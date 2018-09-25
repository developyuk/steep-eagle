import os
from my_eve import schema

MONGO_JSON = 'steep-eagle-3c14071c3df0.json'
ITEM_URL = 'regex("[0-9]{16}")'
OPTIMIZE_PAGINATION_FOR_SPEED = True
PAGINATION_LIMIT = 9999

RESOURCE_METHODS = [
    'GET',
    'POST',
    'DELETE'
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

DOMAIN = {
    'users': {
        'schema': schema.get('users'),
        'soft_delete': True
    },
    'branches': {
        'schema': schema.get('branches')
    },
    'modules': {
        'schema': schema.get('modules')
    },
    'classes': {
        'schema': schema.get('classes')
    },
    'class-students': {
        'schema': schema.get('class-students'),
        'internal_resource': True,
        'disable_documentation': True,
    },
    'classes-students': {
        'url': 'classes/<regex("[0-9]{16}"):class>/students',
        'datasource': {
            'source': 'class-students',
        },
        'schema': schema.get('class-students')
    },
    'attendances': {
        'schema': schema.get('attendances')
    },
    'attendance-tutors': {
        'internal_resource': True,
        'disable_documentation': True,
        'schema': schema.get('attendance-tutors')
    },
    'attendances-tutors': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/tutors',
        'datasource': {
            'source': 'attendance-tutors',
        },
        'schema': schema.get('attendance-tutors')
    },
    'attendance-students': {
        'internal_resource': True,
        'disable_documentation': True,
        'schema': schema.get('attendance-students')
    },
    'attendances-students': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/students',
        'datasource': {
            'source': 'attendance-students',
        },
        'schema': schema.get('attendance-students')
    },
    'attendances-tutors-students': {
        'url': 'attendances/<regex("[0-9]{16}"):attendance>/tutor/<regex("[0-9]{16}"):tutor>/students',
        'datasource': {
            'source': 'attendance-students',
        },
        'schema': schema.get('attendance-students')
    },
    'caches': {
        'schema': schema.get('caches')
    },
}
