

from uuid import uuid4
from pprint import pprint
import os

import jwt
from google.cloud import storage
from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from eve.io.media import MediaStorage
from eve.utils import config


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
        is_exist = len(user)

        if is_exist:
            self.set_request_auth_value(user[config.ID_FIELD])
        # pprint(app.auth.get_request_auth_value())

        return is_exist


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

        'student': {
            'type': 'integer',
            'data_relation': {
                'resource': 'users',
                'field': '_id',
                'embeddable': True
            },
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
        'start': {
            'type': 'datetime'
        },
        'finish': {
            'type': 'datetime'
        },
        'schedule': {
            'type': 'dict',
            'schema': {
                'recurrence': {
                    'type': 'dict',
                    'schema': {
                        'interval': {
                            'type': 'integer'
                        },
                        'freq': {
                            'type': 'string',
                            'allowed': ['daily', 'weekly', 'monthly', 'yearly']
                        },
                        'byday': {
                            'type': 'list',
                            'allowed': ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday']
                        },
                        'until': {
                            'type': 'datetime'
                        },
                        'count': {
                            'type': 'integer'
                        }
                    }
                },
                'exclude': {
                    'type': 'list',
                },
                'include': {
                    'type': 'list',
                },
            }
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
    'schedules': {
        'class': {
            'type': 'integer',
            'required': True,
            'data_relation': {
                'resource': 'classes',
                'field': '_id',
                'embeddable': True
            },
        },
        'start': {
            'type': 'datetime',
            'required': True
        },
        'end': {
            'type': 'datetime',
            'required': True
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
            'data_relation': {
                'resource': 'attendances',
                'field': '_id',
                'embeddable': True
            },
        },
        'tutor': {
            'type': 'integer',
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
        'isPresence': {
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
