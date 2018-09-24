import os
from my_eve import schema

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
# X_ALLOW_CREDENTIALS = True
X_HEADERS = ['Authorization', 'Content-Type',
             'If-Match', 'If-None-Match', 'Cache-Control']
DEBUG = os.environ['DEBUG']
JWT_SECRET = os.environ['JWT_SECRET']
JWT_ISSUER = os.environ['JWT_ISSUER']
X_DOMAINS = "*"

OPTIMIZE_PAGINATION_FOR_SPEED = True

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
        'schema': schema.get('users')
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
    # 'classes-students': {
    #     'url': 'classes/<regex("[a-f0-9]{24}"):class_id>/students',
    #     'datasource': {
    #         'source': 'users'
    #     },
    #     'schema': schema.get('users')

    # },
    # 'attendances': {
    #     'schema': schema.get('attendances')
    # },
}
