import os
import pprint

from eve_sqlalchemy.config import DomainConfig, ResourceConfig
from tables import Users, Branches, Classes, ClassStudents, Modules, Sessions, SessionsTutors, SessionsTutorsStudents, ClassesTs

# from eve_auth_jwt import JWTAuth
SQLALCHEMY_DATABASE_URI = os.environ['SQLALCHEMY_DATABASE_URI']
SQLALCHEMY_TRACK_MODIFICATIONS = False

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
X_DOMAINS = "*"
CACHE_CONTROL = 'max-age=10,must-revalidate'
CACHE_EXPIRES = 10
# X_ALLOW_CREDENTIALS = True
X_HEADERS = ['Authorization', 'Content-Type',
             'If-Match', 'If-None-Match', 'Cache-Control']

DEBUG = os.environ['DEBUG']
JWT_SECRET = os.environ['JWT_SECRET']
JWT_ISSUER = os.environ['JWT_ISSUER']
PAGINATION_DEFAULT = 8
JSONIFY_PRETTYPRINT_REGULAR = False
SWAGGER_INFO = {
    'title': 'M Codingcamp API',
    'version': '1.0',
    'description': 'For https://mtutor.codingcamp.id and https://mops.codingcamp.id',
    'contact': {
        'name': 'developyuk',
        'url': 'https://www.developyuk.com/'
    },
    'schemes': ['http', 'https'],
}
ENFORCE_IF_MATCH = True
UPLOAD_FOLDER = '/tmp/'
PAGINATION_LIMIT = 9999

# The following two lines will output the SQL statements executed by
# SQLAlchemy. This is useful while debugging and in development, but is turned
# off by default.
# --------
# SQLALCHEMY_ECHO = True
# SQLALCHEMY_RECORD_QUERIES = True

# The default schema is generated using DomainConfig:
DOMAIN = DomainConfig({
    'users': ResourceConfig(Users),
    'branches': ResourceConfig(Branches),
    'classes': ResourceConfig(Classes),
    'class_students': ResourceConfig(ClassStudents),
    'modules': ResourceConfig(Modules),
    'sessions': ResourceConfig(Sessions),
    'sessions_tutors': ResourceConfig(SessionsTutors),
    'sessions_tutors_students': ResourceConfig(SessionsTutorsStudents),
    'classes_ts': ResourceConfig(ClassesTs),
}).render()

DOMAIN['modules']['schema']['image'].update({'coerce': 'upload'})
DOMAIN['classes'].update({'allow_unknown': True})
DOMAIN['classes']['schema']['branch']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes']['schema']['tutor']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes']['schema']['students']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes']['schema']['module']['data_relation'].update(
    {'embeddable': True})

DOMAIN['classes_ts']['schema']['branch']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes_ts']['schema']['tutor']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes_ts']['schema']['students']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['classes_ts']['schema']['module']['data_relation'].update(
    {'embeddable': True})

DOMAIN['class_students']['schema']['student']['data_relation'].update(
    {'embeddable': True})
DOMAIN['class_students']['schema']['class_']['data_relation'].update(
    {'embeddable': True})
DOMAIN['sessions']['schema']['session_tutors']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['sessions']['schema']['class_']['data_relation'].update(
    {'embeddable': True})

DOMAIN['sessions_tutors']['schema']['tutor']['data_relation'].update(
    {'embeddable': True})
DOMAIN['sessions_tutors']['schema']['session_students']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['sessions_tutors']['schema']['session']['data_relation'].update(
    {'embeddable': True})

DOMAIN['sessions_tutors_students']['schema']['student']['data_relation'].update(
    {'embeddable': True})
DOMAIN['sessions_tutors_students']['schema']['session_tutor']['data_relation'].update(
    {'embeddable': True})

# pprint.pprint(DOMAIN['modules'])
