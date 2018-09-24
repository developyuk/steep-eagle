

from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from eve.io.base import DataLayer
import jwt
from google.cloud import datastore
from pprint import pprint


def auth(token):
    try:
        data = jwt.decode(token, app.config['JWT_SECRET'])
    except Exception as e:
        abort(400, description=str(e))

    return data


class MyAuth(TokenAuth):

    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials', response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        user = auth(token)
        print(user)

        if len(user):
            self.set_request_auth_value(user[app.config['ID_FIELD']])
        # pprint(app.auth.get_request_auth_value())

        return len(user)


class GoogleCloudstore(DataLayer):
    def init_app(self, app):
        self.driver = datastore.Client.from_service_account_json(
            'steep-eagle-3c14071c3df0.json')

    def find(self, resource, req, sub_resource_lookup):
        args = dict()
        documents = ()

        if req and req.max_results:
            args['limit'] = req.max_results
        if req and req.page > 1:
            args['offset'] = (req.page - 1) * req.max_results

        cursors = self.driver.query(kind=resource).fetch(**args)

        def cursor2list(v):
            v.update({app.config['ID_FIELD']: v.id})
            return dict(v)
        cursors = map(cursor2list, cursors)
        cursors = list(cursors)

        if cursors:
            documents = cursors

        return documents

    def find_one(self, resource, req, check_auth_value=True,
                 force_auth_field_projection=False, **lookup):
        document = dict()

        key = self.driver.key(resource)
        document = self.driver.get(key)

        return document

    def insert(self, resource, doc_or_docs):
        if isinstance(doc_or_docs, dict):
            doc_or_docs = [doc_or_docs]

        incomplete_key = self.driver.key(resource)
        allocate_ids = self.driver.allocate_ids(
            incomplete_key, len(doc_or_docs))

        entities = []
        for i, v in enumerate(doc_or_docs):
            entity = datastore.Entity(allocate_ids[i])
            entity.update(v)
            entities.append(entity)

        self.driver.put_multi(entities)

        return [v.id for v in allocate_ids]


schema = {
    'users': {
        'username': {
            'type': 'string',
        },
        'password': {
            'type': 'string',
        },
        'role': {
            'type': 'string',
            'allowed': ['tutor', 'operation', 'guardian']
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
        # 'photo': {
        #     'type': 'media',
        # },
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
    'branches': {
        'name': {
            'type': 'string',
        },
        # 'address': {
        #     'type': 'string',
        # },
    },
    'modules': {
        'name': {
            'type': 'string',
        },
        # 'image': {
        #     'type': 'media',
        # },
    },
    'classes': {
        'day': {
            'type': 'string',
            'allowed': ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday']
        },
        'startAt': {
            'type': 'string',
        },
        'finishAt': {
            'type': 'string',
        },
        'module': {
            'type': 'string',
        },
        'branch': {
            'type': 'string',
        },
        'tutor': {
            'type': 'string',
        },
    },
    'attendances': {
        'class': {
            'type': 'string',
        },
        'module': {
            'type': 'string',
        },
    }
}
