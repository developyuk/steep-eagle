
from copy import deepcopy
from pprint import pprint
import ast
import json
import dateutil.parser

from google.cloud import datastore
from flask import current_app as app, Response, abort
from eve.utils import config, validate_filters, debug_error_message
from eve.io.mongo.mongo import Mongo
from eve.io.mongo.parser import parse, ParseError
from eve.io.mongo.validation import Validator
import pymongo
from werkzeug.exceptions import HTTPException


class GoogleCloudstoreValidator(Validator):
    def _is_value_unique(self, unique, field, value, query):
        """ Validates that a field value is unique.
        """
        if unique:
            query[field] = value

            resource_config = config.DOMAIN[self.resource]

            # exclude soft deleted documents if applicable
            if resource_config['soft_delete']:
                # be aware that, should a previously (soft) deleted document be
                # restored, and because we explicitly ignore soft deleted
                # documents while validating 'unique' fields, there is a chance
                # that a unique field value will end up being now duplicated
                # in two documents: the restored one, and the one which has
                # been stored with the same field value while the original
                # document was in 'deleted' state.

                # we make sure to also include documents which are missing the
                # DELETED field. This happens when soft deletes are enabled on
                # an a resource with existing documents.
                query[config.DELETED] = {'$ne': True}

            # exclude current document
            if self.document_id:
                id_field = resource_config['id_field']
                query[id_field] = {'$ne': self.document_id}

            # we perform the check on the native mongo driver (and not on
            # app.data.find_one()) because in this case we don't want the usual
            # (for eve) query injection to interfere with this validation. We
            # are still operating within eve's mongo namespace anyway.

            datasource, _, _, _ = app.data.datasource(self.resource)

            if app.data._find_one(self.resource, datasource, query):
                self._error(field, "value '%s' is not unique" % value)


class GoogleCloudstore(Mongo):

    def init_app(self, app):
        driver = datastore.Client.from_service_account_json(
            app.config['MONGO_JSON'])
        self.driver = {'MONGO': driver}

    def set_id_field(self, item):
        if item and not item.get(config.ID_FIELD):
            item.update({config.ID_FIELD: item.id})
        return item

    def conv_projection(self, projections):
        data = []
        for k, v in projections.items():
            if v == 1 and k != config.ID_FIELD:
                data.append(k)
        return tuple(data)

    def conv_order(self, orders):
        data = []
        if orders:
            for v in orders:
                d = v[0]
                if v[1] == -1:
                    d = '-%s' % d
                data.append(d)
        return tuple(data)

    def add_filter(self, query, filter_):
        def _get_equal(k, v):
            return {
                'property_name': k,
                'operator': '=',
                'value': v,
            }

        if filter_:
            for k, v in filter_.items():
                if isinstance(v, dict):
                    for k2, v2 in v.items():
                        f = _get_equal(k, v)

                        try:
                            v2 = int(v2)
                        except (ValueError, TypeError):
                            pass

                        if k == config.DATE_CREATED and isinstance(v2, str):
                            v2 = dateutil.parser.parse(v2)

                        if k2 == '$ne':
                            f = {
                                'property_name': k,
                                'operator': '<',
                                'value': v2,
                            }
                        if k2 == '$ne' and k == config.DELETED:
                            f = {
                                'property_name': k,
                                'operator': '=',
                                'value': not v2,
                            }
                        if k2 == '$gte':
                            f = {
                                'property_name': k,
                                'operator': '>=',
                                'value': v2,
                            }
                        query.add_filter(**f)
                else:
                    try:
                        v = int(v)
                    except ValueError:
                        pass
                    f = _get_equal(k, v)
                    query.add_filter(**f)
        # pprint(query.filters)

    def _find_one(self, resource, datasource, filter_, projection=None):
        args = {
            'filter': filter_,
            'projection': projection,
        }
        document = None
        documents = self._find(resource, datasource, **args)

        if not len(documents):
            return document

        for d in documents:
            document = d
            break

        return document

    def _find(self, resource, datasource, **args):
        filter_ = args.get('filter')
        # pprint(f'resource={resource}')
        # pprint(f'datasource={datasource}')
        # pprint(f'filter={filter_}')

        documents = []

        if filter_:
            values = set()
            for k, v in filter_.items():
                if isinstance(v, list):
                    data = []
                    for v2 in v:
                        v2_args = {
                            'filter': v2
                        }
                        d = self._find(resource, datasource, **v2_args)
                        # pprint(f'd={len(d)}')
                        data.append(d)

                    # pprint('data')
                    # pprint(data)
                    data_id_field = ({d[config.ID_FIELD]
                                      for d in seq} for seq in data)

                    # pprint('data_id_field')
                    # pprint(list(map(list,data_id_field)))

                    if k == '$and':
                        values = set.intersection(*data_id_field)
                    elif k == '$or':
                        values = set.union(*data_id_field)

                    # pprint(f"values={values}")
                    filters = [{config.ID_FIELD: v} for v in values]
                    offset = args.get('skip')
                    offset = offset if offset else 0
                    limit = args.get('limit')
                    limit = limit if limit else 5
                    filters = tuple(filters)[offset:offset+limit]
                    # pprint(f"filters={filters}")

                    if not len(filters):
                        return []

                    keys = map(lambda v: self.pymongo(resource).key(
                        datasource, v[config.ID_FIELD]), filters)

                    documents = self.pymongo(resource).get_multi(list(keys))
                    documents = map(self.set_id_field, documents)
                    documents = list(documents)

                    return documents

        id_field = None
        if filter_ and filter_.get(config.ID_FIELD) and len(filter_.items()) == 1:
            id_field = filter_.get(config.ID_FIELD)
            id_field = int(id_field)

        if id_field:
            key = self.pymongo(resource).key(datasource, id_field)
            document = self.pymongo(resource).get(key)

            documents = [document]
        else:
            order = self.conv_order(args.get('sort'))
            # projection = self.conv_projection(args.get('projection'))
            # pprint(order)
            # pprint(projection)

            query = self.pymongo(resource).query(
                kind=datasource, order=order)

            self.add_filter(query, filter_)
            # pprint(query.filters)

            documents = query.fetch(limit=args.get(
                'limit'), offset=args.get('skip'))

        documents = map(self.set_id_field, documents)
        documents = list(documents)
        return documents

    def allocate_ids(self, datasource, num_ids):
        incomplete_key = self.pymongo(datasource).key(datasource)

        return self.pymongo(datasource).allocate_ids(incomplete_key, num_ids)

    def _insert_many(self, resource, datasource, doc_or_docs, ordered):
        allocate_ids = self.allocate_ids(datasource, len(doc_or_docs))

        entities = []
        for i, v in enumerate(doc_or_docs):
            entity = datastore.Entity(allocate_ids[i])
            entity.update(v)
            entities.append(entity)

        self.pymongo(datasource).put_multi(entities)

        return [v.id for v in allocate_ids]

    def _replace_one(self, resource, datasource, filter_, changes, original):
        id_field = config.DOMAIN[resource]['id_field']
        id_ = filter_[id_field]

        key = self.pymongo(datasource).key(datasource, id_)
        entity = self.pymongo(datasource).get(key)

        original.update(changes)
        entity.update(original)

        self.pymongo(datasource).put(entity)

    def _update_one(self, resource, datasource, filter_, changes, original):
        id_field = config.DOMAIN[resource]['id_field']
        id_ = filter_[id_field]

        key = self.pymongo(datasource).key(datasource, id_)
        entity = self.pymongo(datasource).get(key)
        # pprint(f'{resource} {datasource} {filter_} {changes}')
        entity.update(changes.get('$set'))

        self.pymongo(datasource).put(entity)

    def _delete_many(self, resource, datasource, filter_):
        args = {
            'filter': filter_
        }
        cursors = self._find(resource, datasource, **args)

        keys = [self.pymongo(datasource).key(
            datasource, v[config.ID_FIELD]) for v in cursors]

        self.pymongo(datasource).delete_multi(keys)

    def find_one(self, resource, req, check_auth_value=True,
                 force_auth_field_projection=False, **lookup):
        """ Retrieves a single document.

        :param resource: resource name.
        :param req: a :class:`ParsedRequest` instance.
        :param **lookup: lookup query.
        """
        self._mongotize(lookup, resource)

        client_projection = self._client_projection(req)

        datasource, filter_, projection, _ = self._datasource_ex(
            resource,
            lookup,
            client_projection,
            check_auth_value=check_auth_value,
            force_auth_field_projection=force_auth_field_projection)

        if (config.DOMAIN[resource]['soft_delete']) and \
                (not req or not req.show_deleted) and \
                (not self.query_contains_field(lookup, config.DELETED)):
            filter_ = self.combine_queries(
                filter_, {config.DELETED: {"$ne": True}})
        # Here, we feed pymongo with `None` if projection is empty.
        return self._find_one(resource, datasource, filter_, projection or None)

        # return self.pymongo(resource).db[datasource] \
        #                              .find_one(filter_, projection or None)

    def find(self, resource, req, sub_resource_lookup):
        """ Retrieves a set of documents matching a given request. Queries can
        be expressed in two different formats: the mongo query syntax, and the
        python syntax. The first kind of query would look like: ::

            ?where={"name": "john doe"}

        while the second would look like: ::

            ?where=name=="john doe"

        The resultset if paginated.

        :param resource: resource name.
        :param req: a :class:`ParsedRequest`instance.
        :param sub_resource_lookup: sub-resource lookup from the endpoint url.

        """
        args = dict()

        if req and req.max_results:
            args['limit'] = req.max_results

        if req and req.page > 1:
            args['skip'] = (req.page - 1) * req.max_results

        # TODO sort syntax should probably be coherent with 'where': either
        # mongo-like # or python-like. Currently accepts only mongo-like sort
        # syntax.

        # TODO should validate on unknown sort fields (mongo driver doesn't
        # return an error)

        client_sort = {}
        spec = {}

        if req and req.sort:
            try:
                # assume it's mongo syntax (ie. ?sort=[("name", 1)])
                client_sort = ast.literal_eval(req.sort)
            except ValueError:
                # it's not mongo so let's see if it's a comma delimited string
                # instead (ie. "?sort=-age, name").
                sort = []
                for sort_arg in [s.strip() for s in req.sort.split(",")]:
                    if sort_arg[0] == "-":
                        sort.append((sort_arg[1:], -1))
                    else:
                        sort.append((sort_arg, 1))
                if len(sort) > 0:
                    client_sort = sort
            except Exception as e:
                self.app.logger.exception(e)
                abort(400, description=debug_error_message(str(e)))

        if req and req.where:
            try:
                spec = self._sanitize(json.loads(req.where))
            except HTTPException as e:
                # _sanitize() is raising an HTTP exception; let it fire.
                raise
            except:
                # couldn't parse as mongo query; give the python parser a shot.
                try:
                    spec = parse(req.where)
                except ParseError:
                    abort(400, description=debug_error_message(
                        'Unable to parse `where` clause'
                    ))

        bad_filter = validate_filters(spec, resource)
        if bad_filter:
            abort(400, bad_filter)

        if sub_resource_lookup:
            spec = self.combine_queries(spec, sub_resource_lookup)

        if config.DOMAIN[resource]['soft_delete'] \
                and not (req and req.show_deleted) \
                and not self.query_contains_field(spec, config.DELETED):
            # Soft delete filtering applied after validate_filters call as
            # querying against the DELETED field must always be allowed when
            # soft_delete is enabled
            spec = self.combine_queries(spec, {config.DELETED: {"$ne": True}})

        spec = self._mongotize(spec, resource)

        client_projection = self._client_projection(req)

        datasource, spec, projection, sort = self._datasource_ex(
            resource,
            spec,
            client_projection,
            client_sort)

        if req and req.if_modified_since:
            spec[config.LAST_UPDATED] = \
                {'$gt': req.if_modified_since}

        if len(spec) > 0:
            args['filter'] = spec

        if sort is not None:
            args['sort'] = sort

        if projection:
            args['projection'] = projection

        return self._find(resource, datasource, **args)
        # return self.pymongo(resource).db[datasource].find(**args)

    def insert(self, resource, doc_or_docs):
        """ Inserts a document into a resource collection.

        """
        datasource, _, _, _ = self._datasource_ex(resource)

        # coll = self.get_collection_with_write_concern(datasource, resource)

        if isinstance(doc_or_docs, dict):
            doc_or_docs = [doc_or_docs]

        try:
            return self._insert_many(resource, datasource, doc_or_docs, ordered=True)
            # return coll.insert_many(doc_or_docs, ordered=True).inserted_ids
        except pymongo.errors.BulkWriteError as e:
            self.app.logger.exception(e)

            # since this is an ordered bulk operation, all remaining inserts
            # are aborted. Be aware that if BULK_ENABLED is True and more than
            # one document is included with the payload, some documents might
            # have been successfully inserted, even if the operation was
            # aborted.

            # report a duplicate key error since this can probably be
            # handled by the client.
            for error in e.details['writeErrors']:
                # amazingly enough, pymongo does not appear to be exposing
                # error codes as constants.
                if error['code'] == 11000:
                    abort(409, description=debug_error_message(
                        'Duplicate key error at index: %s, message: %s' % (
                            error['index'], error['errmsg'])
                    ))

            abort(500, description=debug_error_message(
                'pymongo.errors.BulkWriteError: %s' % e
            ))

    def _change_request(self, resource, id_, changes, original, replace=False):
        """ Performs a change, be it a replace or update.

        .. versionchanged:: 0.6.1
           Support for PyMongo 3.0.

        .. versionchanged:: 0.6
           Return 400 if an attempt is made to update/replace an immutable
           field.
        """
        id_field = config.DOMAIN[resource]['id_field']
        query = {id_field: id_}
        if config.ETAG in original:
            query[config.ETAG] = original[config.ETAG]

        datasource, filter_, _, _ = self._datasource_ex(
            resource, query)

        # coll = self.get_collection_with_write_concern(datasource, resource)
        try:
            self._replace_one(resource, datasource, filter_, changes, original) if replace else \
                self._update_one(resource, datasource,
                                 filter_, changes, original)
            # coll.replace_one(filter_, changes) if replace else \
            #     coll.update_one(filter_, changes)
        except pymongo.errors.DuplicateKeyError as e:
            abort(400, description=debug_error_message(
                'pymongo.errors.DuplicateKeyError: %s' % e
            ))
        except pymongo.errors.OperationFailure as e:
            # server error codes and messages changed between 2.4 and 2.6/3.0.
            server_version = \
                self.driver.db.client.server_info()['version'][:3]
            if (
                (server_version == '2.4' and e.code in (13596, 10148)) or
                (server_version in ('2.6', '3.0', '3.2', '3.4') and
                    e.code in (66, 16837))
            ):
                # attempt to update an immutable field. this usually
                # happens when a PATCH or PUT includes a mismatching ID_FIELD.
                self.app.logger.warning(e)
                description = debug_error_message(
                    'pymongo.errors.OperationFailure: %s' % e) or \
                    "Attempt to update an immutable field. Usually happens " \
                    "when PATCH or PUT include a '%s' field, " \
                    "which is immutable (PUT can include it as long as " \
                    "it is unchanged)." % id_field

                abort(400, description=description)
            else:
                # see comment in :func:`insert()`.
                self.app.logger.exception(e)
                abort(500, description=debug_error_message(
                    'pymongo.errors.OperationFailure: %s' % e
                ))

    def find_one_raw(self, resource, **lookup):
        """ Retrieves a single raw document.

        :param resource: resource name.
        :param **lookup: lookup query.

        """
        id_field = config.DOMAIN[resource]['id_field']
        _id = lookup.get(id_field)
        datasource, filter_, _, _ = self._datasource_ex(resource,
                                                        {id_field: _id},
                                                        None)

        lookup = self._mongotize(lookup, resource)

        return self._find_one(resource, datasource, lookup)
        # return self.pymongo(resource).db[datasource].find_one(lookup)

    def remove(self, resource, lookup):
        """ Removes a document or the entire set of documents from a
        collection.

        :returns
            A document (dict) describing the effect of the remove
            or None if write acknowledgement is disabled.
        """
        lookup = self._mongotize(lookup, resource)
        datasource, filter_, _, _ = self._datasource_ex(resource, lookup)

        # coll = self.get_collection_with_write_concern(datasource, resource)
        try:
            self._delete_many(resource, datasource, filter_)
            # coll.delete_many(filter_)
        except pymongo.errors.OperationFailure as e:
            # see comment in :func:`insert()`.
            self.app.logger.exception(e)
            abort(500, description=debug_error_message(
                'pymongo.errors.OperationFailure: %s' % e
            ))
