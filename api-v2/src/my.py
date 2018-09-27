

from uuid import uuid4
from pprint import pprint
import json
import ast
import os
from copy import deepcopy

import jwt
from google.cloud import datastore, storage
from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from eve.io.media import MediaStorage
from eve.io.mongo.mongo import Mongo
from eve.io.mongo.parser import parse, ParseError
from eve.utils import config, validate_filters, debug_error_message
from werkzeug.exceptions import HTTPException


def auth(token):
    try:
        data = jwt.decode(token, config.JWT_SECRET)
    except Exception as e:
        abort(400, description=str(e))

    return data


class JwtAuth(TokenAuth):
    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials', response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        user = auth(token)

        if len(user):
            self.set_request_auth_value(user[config.ID_FIELD])
        # pprint(app.auth.get_request_auth_value())

        return len(user)


class GoogleMediaStorage(MediaStorage):
    def put(self, content, filename=None, content_type=None, resource=None):
        if resource in ('modules', 'users'):
            ext = filename.rsplit('.', 1)[1].lower()
            name = uuid4().hex+'.'+ext
            # name = secure_filename(content.filename)
            location = os.path.join(app.config['UPLOAD_FOLDER'], name)
            content.save(location)
            client = storage.Client.from_service_account_json(
                app.config['MONGO_JSON'])
            bucket = client.bucket(app.config['MEDIA_ENDPOINT'])
            blob = bucket.blob(name)

            blob.upload_from_filename(location, content_type=content_type)

            return name

    def get(self, id_or_filename, resource=None):
        if resource in ('modules', 'users'):
            return id_or_filename

    def delete(self, id_or_filename, resource=None):
        pass


class GoogleCloudstore(Mongo):

    def init_app(self, app):
        driver = datastore.Client.from_service_account_json(
            app.config['MONGO_JSON'])
        self.driver = {'MONGO': driver}

    def remap_entity(self, v):
        if not v.get(config.ID_FIELD):
            v.update({config.ID_FIELD: v.id})
        return v

    def convert_sort(self, v):
        return v[0] if v[1] == 1 else '-%s' % v[0]

    def combine_queries(self, query_a, query_b):
        query = deepcopy(query_a)
        query.update(query_b)

        return query

    def convert_queries(self, _queries):
        queries = []
        for k, v in _queries.items():
            op = '='
            if isinstance(v, dict):
                for k2, v2 in v.items():
                    if k2 == '$ne' and k == '_deleted' and v2:
                        op = '<'
                    if k2 == '$ne' and k == '_deleted' and not v2:
                        op = '>'

                    q = {
                        'property_name': k,
                        'operator': op,
                        'value': v2
                    }
                    queries.append(q)

            elif isinstance(v, list):
                for v2 in v:
                    for k3, v3 in v2.items():
                        if k == '$or' and k3 == '_id':
                            k = k3
                        q = {
                            'property_name': k,
                            'operator': op,
                            'value': v3
                        }
                        queries.append(q)
            else:
                q = {
                    'property_name': k,
                    'operator': op,
                    'value': v
                }
                queries.append(q)
        return queries

    def find(self, resource, req, sub_resource_lookup):
        args = dict()
        documents = ()

        if req and req.max_results:
            args['limit'] = req.max_results

        if req and req.page > 1:
            args['offset'] = (req.page - 1) * req.max_results

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

        query = self.pymongo(datasource).query(kind=datasource)

        # pprint(datasource)
        # pprint(spec)
        index_id_field = -1
        if len(spec) > 0:
            _spec = self.convert_queries(spec)

            for k, v in enumerate(_spec):
                if v['property_name'] == '_id':
                    index_id_field = k

        if index_id_field >= 0:
            key = self.pymongo(datasource).key(
                datasource, int(_spec[index_id_field]['value']))
            cursors = self.pymongo(resource).get(key)

            if cursors:
                # pprint(_filter_)
                for v in _spec:
                    if v['property_name'] != '_id' and v['operator'] != '=' and cursors[v['property_name']] == v['value']:
                        cursors = None
            if cursors:
                cursors = self.remap_entity(cursors)
                cursors = [cursors]
        else:
            _spec = self.convert_queries(spec)

            for v in _spec:
                query.add_filter(**v)

            if sort is not None:
                sort = map(self.convert_sort, sort)
                query.order = tuple(sort)

            # if client_projection:
            #     projection = list(projection.keys())
            #     projection.remove('_id')

            #     for v in spec.keys():
            #         try:
            #             projection.remove(v)
            #         except ValueError:
            #             pass
            #     query.projection = tuple(projection)

            cursors = query.fetch(**args)

            # pprint(query.filters)
            # pprint(query.order)
            cursors = map(self.remap_entity, cursors)
            cursors = list(cursors)

        if cursors:
            documents = cursors

        return documents

    def find_one(self, resource, req, check_auth_value=True,
                 force_auth_field_projection=False, **lookup):
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

        filter_id_field = filter_.get(config.ID_FIELD)
        cursors = None
        # pprint(datasource)
        # pprint(filter_)
        if filter_id_field:
            key = self.pymongo(datasource).key(
                datasource, int(filter_id_field))
            cursors = self.pymongo(resource).get(key)

            if cursors:
                _filter_ = self.convert_queries(filter_)
                # pprint(_filter_)
                for v in _filter_:
                    if v['property_name'] != '_id' and v['operator'] != '=' and cursors[v['property_name']] == v['value']:
                        cursors = None
            if cursors:
                cursors = self.remap_entity(cursors)
        else:
            query = self.pymongo(datasource).query(kind=datasource)

            _filter_ = self.convert_queries(filter_)
            for v in _filter_:
                query.add_filter(**v)
            cursors = query.fetch(limit=1)

            cursors = map(self.remap_entity, cursors)
            cursors = list(cursors)

            if cursors:
                cursors = cursors[0]

            # pprint(query.filters)
            # pprint(query.order)
            # pprint(cursors)
        return cursors

    def find_one_raw(self, resource, **lookup):

        id_field = config.DOMAIN[resource]['id_field']
        _id = lookup.get(id_field)
        datasource, filter_, _, _ = self._datasource_ex(resource,
                                                        {id_field: _id},
                                                        None)

        lookup = self._mongotize(lookup, resource)

        filter_id_field = filter_.get(config.ID_FIELD)
        cursors = None
        if filter_id_field:
            key = self.pymongo(datasource).key(
                datasource, int(filter_id_field))
            cursors = self.pymongo(resource).get(key)

            if not cursors:
                abort(404, description='item not found')

            cursors = self.remap_entity(cursors)
        else:
            pass

        return cursors
        # return self.pymongo(resource).db[datasource].find_one(lookup)

    def insert(self, resource, doc_or_docs):
        datasource, _, _, _ = self._datasource_ex(resource)

        if isinstance(doc_or_docs, dict):
            doc_or_docs = [doc_or_docs]

        incomplete_key = self.pymongo(datasource).key(datasource)
        allocate_ids = self.pymongo(datasource).allocate_ids(
            incomplete_key, len(doc_or_docs))

        entities = []
        for i, v in enumerate(doc_or_docs):
            entity = datastore.Entity(allocate_ids[i])
            entity.update(v)
            entities.append(entity)

        self.pymongo(datasource).put_multi(entities)

        return [v.id for v in allocate_ids]

    def _change_request(self, resource, id_, changes, original, replace=False):
        id_field = config.DOMAIN[resource]['id_field']
        query = {id_field: id_}
        if config.ETAG in original:
            query[config.ETAG] = original[config.ETAG]

        datasource, filter_, _, _ = self._datasource_ex(
            resource, query)

        key = self.pymongo(datasource).key(datasource, id_)
        entity = self.pymongo(datasource).get(key)
        if replace:
            original.update(changes.get('$set'))
            entity.update(original)
        else:
            entity.update(changes.get('$set'))

        self.pymongo(datasource).put(entity)

    def remove(self, resource, lookup):
        lookup = self._mongotize(lookup, resource)
        datasource, filter_, _, _ = self._datasource_ex(resource, lookup)

        filter_id_field = filter_.get(config.ID_FIELD)
        if filter_id_field:
            key = self.pymongo(datasource).key(
                datasource, int(filter_id_field))
            cursors = self.pymongo(resource).get(key)

            if not cursors:
                abort(404, description='item not found')

            cursors = self.remap_entity(cursors)
            cursors = [cursors]
        else:
            query = self.pymongo(datasource).query(kind=datasource)
            for k, v in filter_.items():
                query.add_filter(k, '=', int(v))
            cursors = query.fetch()

            cursors = map(self.remap_entity, cursors)
            cursors = list(cursors)

        keys = [self.pymongo(datasource).key(
            datasource, v[config.ID_FIELD]) for v in cursors]

        self.pymongo(datasource).delete_multi(keys)


schema = {
    'user': {
        'username': {
            'type': 'string'
        },
        'password': {
            'type': 'string',
        },
        'role': {
            'type': 'string',
            'required': True,
            'allowed': ['tutor', 'operation', 'student', 'guardian']
        },
        'email': {
            'type': 'string',
        },
        'name': {
            'type': 'string',
        },
        'dob': {
            'type': 'datetime',
        },
        'photo': {
            'type': 'media',
        },
        'address': {
            'type': 'string',
        },
        'school': {
            'type': 'string',
        },
        'grade': {
            'type': 'string',
        },
        'contact': {
            'type': 'string',
        },
    },
    'branch': {
        'name': {
            'type': 'string',
            'required': True
        },
        'address': {
            'type': 'string',
        },
    },
    'module': {
        'name': {
            'type': 'string',
            'required': True
        },
        'image': {
            'type': 'media',
        },
    },
    'class': {
        'day': {
            'type': 'string',
            'required': True,
            'allowed': ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday']
        },
        'startAt': {
            'type': 'string',
            'required': True
        },
        'finishAt': {
            'type': 'string',
            'required': True
        },
        'module': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'modules',
                'field': '_id',
                'embeddable': True
            },
        },
        'branch': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'branches',
                'field': '_id',
                'embeddable': True
            },
        },
        'tutor': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
        },
    },
    'class-student': {
        'class': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'classes',
                'field': '_id',
                'embeddable': True
            },
        },
        'student': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
        },
    },
    'attendance': {
        'class': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'classes',
                'field': '_id',
                'embeddable': True
            },
        },
        'module': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'modules',
                'field': '_id',
                'embeddable': True
            },
        },
    },
    'attendance-tutor': {
        'attendance': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'attendances',
                'field': '_id',
                'embeddable': True
            },
        },
        'tutor': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
        },
    },
    'attendance-student': {
        'attendance': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'attendances',
                'field': '_id',
                'embeddable': True
            },
        },
        'tutor': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
        },
        'student': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
        },
        'is_presence': {
            'type': 'boolean',
            'default': False
        },
        'feedback': {
            'type': 'string',
        },
        'rating': {
            'type': 'dict',
            'schema': {
                'interaction': {'type': 'integer'},
                'cognition': {'type': 'integer'},
                'creativity': {'type': 'integer'},
            },

        },
    },
    'cache': {
        'key': {
            'type': 'string',
        },
        'value': {
            'type': 'dict',
        },
    }
}
