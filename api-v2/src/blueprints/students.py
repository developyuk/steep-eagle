from datetime import timedelta
from copy import deepcopy
from pprint import pprint

from flask_cors import CORS
from eve_swagger import add_documentation
from flask import current_app as app, jsonify, Blueprint
from eve.auth import requires_auth
from eve.methods.get import get_internal, getitem_internal
from eve.methods.put import put_internal
from eve.methods.post import post_internal
from eve.render import send_response
from eve.utils import config

from shared.datetime import utc_now, after_request_cache, wib_tz

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(seconds=config.CACHE_EXPIRES))


def not_in_attendance_students(attendances_students, attendances_tutors, item):
    attendances_tutors_created = attendances_tutors[config.DATE_CREATED].astimezone(wib_tz).replace(
        hour=0, minute=0, second=0, microsecond=0)

    for v in attendances_students:
        if (v['student'] == item['student'] and
                v['attendance'] == attendances_tutors['attendance'][config.ID_FIELD] and
                v[config.DATE_CREATED] >= attendances_tutors_created):
            return False

    return True


add_documentation({
    'paths': {'/students/attendances': {'get': {
        'summary': 'Retrieves one or more student attendance',
        "responses": {
            "200": {
                "description": "Student document fetched successfully",
                "schema": {"$ref": "#/definitions/User"}
            }
        },
        'tags': ['User'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/students/attendances', methods=['GET'])
@requires_auth('/students/attendances')
def get_students():

    attendances, *_ = get_internal('attendances')
    attendances = attendances[config.ITEMS]

    attendances_students, *_ = get_internal('attendances_students')
    attendances_students = attendances_students[config.ITEMS]

    lookup = {
        config.DATE_CREATED: {'$gte': (utc_now - timedelta(hours=12))},
        'tutor': app.auth.get_request_auth_value()
    }
    attendances_tutors, *_ = get_internal('attendances_tutors', **lookup)
    attendances_tutors = attendances_tutors[config.ITEMS]

    for v in attendances_tutors:
        for v2 in attendances:
            if v['attendance'] == v2[config.ID_FIELD]:
                v['attendance'] = v2

    for v in attendances_tutors:
        lookup = {
            'class': v['attendance']['class']
        }
        students, *_ = get_internal('classes_students', **lookup)
        students = students[config.ITEMS]

        students = filter(
            lambda v2: not_in_attendance_students(attendances_students, v, v2), students)
        v['students'] = list(students)

    attendances_tutors = filter(lambda v: len(
        v['students']) > 0, attendances_tutors)
    attendances_tutors = {config.ITEMS: list(attendances_tutors)}

    return jsonify(attendances_tutors)


def _dormant_students(classes_students, student):
    for v in classes_students:
        if v['student'] == student[config.ID_FIELD]:
            return False
    return True


add_documentation({
    'paths': {'/students/dormant': {'get': {
        'summary': 'Retrieves one or more dormant student',
        "responses": {
            "200": {
                "description": "Student document fetched successfully",
                "schema": {"$ref": "#/definitions/User"}
            }
        },
        'tags': ['User'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/students/dormant', methods=['GET'])
@requires_auth('/students/dormant')
def students_dormant():
    resource = 'users'
    lookup = {
        'role': 'student'
    }
    response = get_internal(resource, **lookup)
    response = list(response)
    students = response[0][config.ITEMS]

    classes_students, *_ = get_internal('classes_students')
    classes_students = classes_students[config.ITEMS]

    students = filter(lambda v: _dormant_students(
        classes_students, v), students)
    response[0][config.ITEMS] = list(students)

    return send_response(resource, response)


@blueprint.after_request
def add_header(response):
    return after_request_cache(response)
