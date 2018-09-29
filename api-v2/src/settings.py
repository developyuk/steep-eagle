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
MEDIA_ENDPOINT = 'steep-eagle-files'
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
        'schema': schema.get('user'),
        'soft_delete': True
    },
    'students': {
        'datasource': {
            'source': 'users',
            'filter': {'role': 'student'}
        },
        'schema': schema.get('user'),
        'soft_delete': True
    },
    'students-guardians': {
        'url': 'students/<regex("[0-9]{16}"):student>/guardians',
        'datasource': {
            'source': 'users',
            'filter': {'role': 'guardian'}
        },
        'schema': schema.get('user'),
        'resource_methods': ['GET', 'POST', 'DELETE']
    },
    'tutors': {
        'datasource': {
            'source': 'users',
            'filter': {'role': 'tutor'}
        },
        'schema': schema.get('user'),
        'soft_delete': True
    },
    'branches': {
        'schema': schema.get('branch')
    },
    'modules': {
        'schema': schema.get('module')
    },
    'classes': {
        'schema': schema.get('class')
    },
    'classes_students': {
        'schema': schema.get('class-student'),
        'internal_resource': True,
        'disable_documentation': True,
    },
    'classes-students': {
        'url': 'classes/<regex("[0-9]{16}"):class>/students',
        'datasource': {
            'source': 'classes_students',
        },
        'schema': schema.get('class-student'),
        'resource_methods': ['GET', 'POST', 'DELETE']
    },
    'attendances': {
        'schema': schema.get('attendance'),
        'item_methods': ['GET']
    },
    'attendances_tutors': {
        'schema': schema.get('attendance-tutor'),
        'item_methods': ['GET'],
        'allow_unknown': True
    },
    'attendances-tutors': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/tutors',
        'datasource': {
            'source': 'attendances_tutors',
        },
        'schema': schema.get('attendance-tutor'),
        'item_methods': ['GET']
    },
    'attendances_students': {
        'schema': schema.get('attendance-student')
    },
    'attendances-students': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/students',
        'datasource': {
            'source': 'attendances_students',
        },
        'schema': schema.get('attendance-student'),
        'item_methods': ['GET']
    },
    'caches': {
        'schema': schema.get('cache'),
        'item_methods': ['GET', 'PATCH']
    },
}
