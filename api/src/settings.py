import os
# from pprint import pprint

from eve_sqlalchemy.config import DomainConfig, ResourceConfig
import tables

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
    'schemes': ['https'],
}
ENFORCE_IF_MATCH = True
PAGINATION_LIMIT = 9999

RETURN_MEDIA_AS_BASE64_STRING = False
RETURN_MEDIA_AS_URL = True
MEDIA_BASE_URL = 'http://%s:9000' % os.environ['HOST_STORAGE']
MEDIA_ENDPOINT = 'module-images'
UPLOAD_FOLDER = '/tmp/'

MAIL_SERVER = 'smtp.gmail.com'
MAIL_PORT = 465
MAIL_USE_TLS = False
MAIL_USE_SSL = True
# MAIL_DEBUG : default app.debug
MAIL_USERNAME = 'andry.yosua@gmail.com'
MAIL_PASSWORD = os.environ['MAIL_PASSWORD']
MAIL_DEFAULT_SENDER = ('no-reply', 'andry.yosua@gmail.com')
# MAIL_MAX_EMAILS : default None
# MAIL_SUPPRESS_SEND : default app.testing
# MAIL_ASCII_ATTACHMENTS : default False

# The following two lines will output the SQL statements executed by
# SQLAlchemy. This is useful while debugging and in development, but is turned
# off by default.
# --------
# SQLALCHEMY_ECHO = True
# SQLALCHEMY_RECORD_QUERIES = True

# The default schema is generated using DomainConfig:
DOMAIN = DomainConfig({
    'users': ResourceConfig(tables.Users),
    'branches': ResourceConfig(tables.Branches),
    'classes': ResourceConfig(tables.Classes),
    'class_students': ResourceConfig(tables.ClassStudents),
    'modules': ResourceConfig(tables.Modules),
    'attendances': ResourceConfig(tables.Attendances),
    'attendances_tutors': ResourceConfig(tables.AttendancesTutors),
    'attendances_students': ResourceConfig(tables.AttendancesStudents),
    'classes_ts': ResourceConfig(tables.ClassesTs),
    'notifications': ResourceConfig(tables.Notifications),
    'student_guardians': ResourceConfig(tables.StudentGuardians),
    'caches': ResourceConfig(tables.Caches),
}).render()

DOMAIN['modules']['schema']['image'].update({'type': 'media'})
DOMAIN['users']['schema']['photo'].update({'type': 'media'})

DOMAIN['classes'].update({'allow_unknown': True})
DOMAIN['attendances_tutors'].update({'allow_unknown': True})
DOMAIN['users'].update({'allow_unknown': True})

# DOMAIN['users'].update({'soft_delete': True})

DOMAIN['users']['schema']['student_guardians']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['student_guardians']['schema']['guardian']['data_relation'].update(
    {'embeddable': True})
# pprint(DOMAIN['users'])
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
DOMAIN['attendances']['schema']['attendance_tutors']['schema']['data_relation'].update(
    {'embeddable': True})
DOMAIN['attendances']['schema']['class_']['data_relation'].update(
    {'embeddable': True})
DOMAIN['attendances']['schema']['module']['data_relation'].update(
    {'embeddable': True})

DOMAIN['attendances_tutors']['schema']['tutor']['data_relation'].update(
    {'embeddable': True})
# DOMAIN['attendances_tutors']['schema']['attendance_students']['schema']['data_relation'].update( {'embeddable': True})
DOMAIN['attendances_tutors']['schema']['attendance']['data_relation'].update(
    {'embeddable': True})

DOMAIN['attendances_students']['schema']['student']['data_relation'].update(
    {'embeddable': True})
DOMAIN['attendances_students']['schema']['tutor']['data_relation'].update(
    {'embeddable': True})
DOMAIN['attendances_students']['schema']['attendance']['data_relation'].update(
    {'embeddable': True})
# DOMAIN['attendances_students']['schema']['attendance_tutor']['data_relation'].update( {'embeddable': True})
