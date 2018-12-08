import os
from my import schema

MONGO_JSON = os.environ['MONGO_JSON']
ITEM_URL = 'regex("[0-9]{16}")'
OPTIMIZE_PAGINATION_FOR_SPEED = True
PAGINATION_LIMIT = 9999

RESOURCE_METHODS = [
    'GET',
    'POST',
    # 'DELETE'
]
ITEM_METHODS = [
    'GET',
    'PATCH',
    # 'PUT',
    'DELETE'
]
CACHE_CONTROL = 'max-age=10,must-revalidate'
CACHE_EXPIRES = 10
X_HEADERS = ['Authorization', 'Content-Type',
             'If-Match', 'If-None-Match', 'Cache-Control']
DEBUG = os.environ['DEBUG']
JWT_SECRET = os.environ['JWT_SECRET']
JWT_ISSUER = os.environ['JWT_ISSUER']
X_DOMAINS = "*"

RETURN_MEDIA_AS_BASE64_STRING = False
RETURN_MEDIA_AS_URL = True
MEDIA_BASE_URL = 'https://storage.googleapis.com'
MEDIA_ENDPOINT = 'm-files'
UPLOAD_FOLDER = '/tmp/'

SWAGGER_INFO = {
    'title': 'M Codingcamp API',
    'version': '1.0',
    'description': 'For https://mtutor.codingcamp.id and https://mops.codingcamp.id',
    'contact': {
        'name': 'developyuk',
        'url': 'https://www.developyuk.com/'
    },
    'schemes': ['http'],
}
RENDERERS = ['eve.render.JSONRenderer']
XML = False

DOMAIN = {
    'users': {
        'schema': schema.user,
        'soft_delete': True
    },
    'students-guardians': {
        'url': 'users/<regex("[0-9]{16}"):student>/guardians',
        'datasource': {
            'source': 'users',
            'filter': {'role': 'guardian'}
        },
        'schema': schema.user,
        'resource_methods': ['GET', 'POST', 'DELETE']
    },
    'branches': {
        'schema': schema.branch
    },
    'modules': {
        'schema': schema.module
    },
    'classes': {
        'schema': schema.class_,
        'allow_unknown': True
    },
    'schedules': {
        'schema': schema.schedules,
    },
    'classes_students': {
        'schema': schema.class_student,
        'internal_resource': True,
        'disable_documentation': True,
    },
    'classes-students': {
        'url': 'classes/<regex("[0-9]{16}"):class>/students',
        'datasource': {
            'source': 'classes_students',
        },
        'schema': schema.class_student,
        'resource_methods': ['GET', 'POST', 'DELETE']
    },
    'attendances': {
        'schema': schema.attendance,
        'item_methods': ['GET']
    },
    'attendances_tutors': {
        'schema': schema.attendance_tutor,
        'item_methods': ['GET', 'DELETE'],
        'allow_unknown': True
    },
    'attendances-tutors': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/tutors',
        'datasource': {
            'source': 'attendances_tutors',
        },
        'schema': schema.attendance_tutor,
        'item_methods': ['GET']
    },
    'attendances_students': {
        'schema': schema.attendance_student
    },
    'attendances-students': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/students',
        'datasource': {
            'source': 'attendances_students',
        },
        'schema': schema.attendance_student,
        'item_methods': ['GET']
    },
    'caches': {
        'schema': schema.cache,
        'item_methods': ['GET', 'PATCH']
    },
}
