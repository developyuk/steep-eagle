import os
from my import schema

MONGO_HOST = 'db'
MONGO_PORT = 27017
# Skip these if your db has no auth. But it really should.
MONGO_USERNAME = os.environ['MONGO_USERNAME']
MONGO_PASSWORD = os.environ['MONGO_PASSWORD']
MONGO_DBNAME = 'admin'
# ITEM_URL = 'regex("[0-9]{16}")'
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
PAGINATION_LIMIT = 9999
OPTIMIZE_PAGINATION_FOR_SPEED = True
CACHE_CONTROL = 'max-age=10,must-revalidate'
CACHE_EXPIRES = 10
X_HEADERS = ['Authorization', 'Content-Type',
             'If-Match', 'If-None-Match', 'Cache-Control']
DEBUG = os.environ['DEBUG']
JWT_SECRET = os.environ['JWT_SECRET']
JWT_ISSUER = os.environ['JWT_ISSUER']
X_DOMAINS = "*"

DATE_FORMAT = '%Y-%m-%d %H:%M:%S'
RETURN_MEDIA_AS_BASE64_STRING = False
RETURN_MEDIA_AS_URL = True
MEDIA_BASE_URL = 'https://storage.googleapis.com'
MEDIA_ENDPOINT = 'steep-eagle-files'
# MEDIA_ENDPOINT = 'm-files'
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

MAIL_SERVER = 'smtp.gmail.com'
MAIL_PORT = 465
MAIL_USE_SSL = True
MAIL_USERNAME = 'andry.yosua@gmail.com'
MAIL_DEFAULT_SENDER = 'andry.yosua@gmail.com'
MAIL_PASSWORD = os.environ['MAIL_PASSWORD']

DOMAIN = {
    'users': {
        'schema': schema.user,
        'soft_delete': True
    },
    'students-guardians': {
        'url': 'users/<regex("[a-f0-9]{24}"):student>/guardians',
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
        'schema': schema.class_
    },
    'classes_students': {
        'schema': schema.class_student,
        'internal_resource': True,
        'disable_documentation': True,
    },
    'classes-students': {
        'url': 'classes/<regex("[a-f0-9]{24}"):class>/students',
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
        'url': 'attendances/<regex("[a-f0-9]{24}"):attendance>/tutors',
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
        'url': 'attendances/<regex("[a-f0-9]{24}"):attendance>/students',
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
