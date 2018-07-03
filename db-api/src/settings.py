from eve_sqlalchemy.config import DomainConfig, ResourceConfig
from tables import Users, UsersProfile, Branches, Classes, ClassStudents, Modules, ProgramTypes, Programs, \
  ProgramsModules, Sessions, SessionsTutors, SessionsTutorsStudents, \
  Exports, ClassesTs
import os
import pprint

# from eve_auth_jwt import JWTAuth
SQLALCHEMY_DATABASE_URI = os.environ['SQLALCHEMY_DATABASE_URI']
SQLALCHEMY_TRACK_MODIFICATIONS = False

RESOURCE_METHODS = [
  'GET',
  # 'POST', 'DELETE'
]

ITEM_METHODS = [
  'GET',
  # 'PATCH', 'PUT', 'DELETE'
]
X_DOMAINS = "*"
CACHE_CONTROL = 'max-age=10,must-revalidate'
CACHE_EXPIRES = 10

DEBUG = os.environ['DEBUG']
JWT_SECRET = os.environ['JWT_SECRET']
JWT_ISSUER = os.environ['JWT_ISSUER']
PAGINATION_DEFAULT = 8
EMBEDDING = True
JSONIFY_PRETTYPRINT_REGULAR = False

# The following two lines will output the SQL statements executed by
# SQLAlchemy. This is useful while debugging and in development, but is turned
# off by default.
# --------
# SQLALCHEMY_ECHO = True
# SQLALCHEMY_RECORD_QUERIES = True

# The default schema is generated using DomainConfig:
DOMAIN = DomainConfig({
  'users': ResourceConfig(Users),
  'users_profile': ResourceConfig(UsersProfile),
  'branches': ResourceConfig(Branches),
  'classes': ResourceConfig(Classes),
  'class_students': ResourceConfig(ClassStudents),
  'modules': ResourceConfig(Modules),
  'program_types': ResourceConfig(ProgramTypes),
  'programs': ResourceConfig(Programs),
  'programs_modules': ResourceConfig(ProgramsModules),
  'sessions': ResourceConfig(Sessions),
  'sessions_tutors': ResourceConfig(SessionsTutors),
  'sessions_tutors_students': ResourceConfig(SessionsTutorsStudents),
  'exports': ResourceConfig(Exports),
  'classes_ts': ResourceConfig(ClassesTs),
}).render()

DOMAIN['users']['schema']['profile']['data_relation'].update({u'embeddable': True})
DOMAIN['users'].update({u'embedded_fields': ['profile']})
DOMAIN['classes']['schema']['branch']['data_relation'].update({u'embeddable': True})
DOMAIN['classes']['schema']['tutor']['data_relation'].update({u'embeddable': True})
DOMAIN['classes']['schema']['students']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['classes']['schema']['program_module']['data_relation'].update({u'embeddable': True})
DOMAIN['classes'].update({u'embedded_fields': [
  # 'branch',
  # 'tutor', 'tutor.profile',
  # 'students', 'students.student', 'students.student.profile',
  # 'program_module', 'program_module.module', 'program_module.program',
  # 'program_module.program.type'
]})
DOMAIN['classes_ts']['schema']['branch']['data_relation'].update({u'embeddable': True})
DOMAIN['classes_ts']['schema']['tutor']['data_relation'].update({u'embeddable': True})
DOMAIN['classes_ts']['schema']['students']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['classes_ts']['schema']['program_module']['data_relation'].update({u'embeddable': True})
DOMAIN['classes_ts'].update({u'embedded_fields': [
  'branch',
  'tutor', 'tutor.profile',
  # 'students', 'students.student', 'students.student.profile',
  'program_module', 'program_module.module', 'program_module.program',
  'program_module.program.type'
]})
DOMAIN['programs_modules']['schema']['module']['data_relation'].update({u'embeddable': True})
DOMAIN['programs_modules']['schema']['program']['data_relation'].update({u'embeddable': True})
DOMAIN['class_students']['schema']['student']['data_relation'].update({u'embeddable': True})
DOMAIN['class_students']['schema']['class_']['data_relation'].update({u'embeddable': True})
DOMAIN['programs']['schema']['type']['data_relation'].update({u'embeddable': True})
DOMAIN['programs']['schema']['modules']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['programs'].update({u'embedded_fields': [
  'type',
  'modules', 'modules.module'
]})
DOMAIN['modules']['schema']['programs']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['modules'].update({u'embedded_fields': [
  'programs', 'programs.program', 'programs.program.type'
]})
DOMAIN['sessions']['schema']['session_tutors']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['sessions'].update({u'embedded_fields': [
  'session_tutors', 'session_tutors.tutor', 'session_tutors.tutor.profile',
  'session_tutors.session_students', 'session_tutors.session_students.student',
  'session_tutors.session_students.student.profile'
]})
DOMAIN['sessions_tutors']['schema']['tutor']['data_relation'].update({u'embeddable': True})
DOMAIN['sessions_tutors']['schema']['session_students']['schema']['data_relation'].update({u'embeddable': True})
DOMAIN['sessions_tutors_students']['schema']['student']['data_relation'].update({u'embeddable': True})
# pprint.pprint(DOMAIN)
