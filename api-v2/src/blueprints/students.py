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

from shared.datetime import utc_now, after_request_cache

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(seconds=10))


def attendance_students(item, at):
    lookup = {
        config.DATE_CREATED: {'$gte': at[config.DATE_CREATED].replace(hour=0, minute=0, second=0, microsecond=0)},
        'student': item['student'],
        'attendance': at['attendance'][config.ID_FIELD],
    }
    attendances_students, *_ = get_internal('attendances_students', **lookup)
    attendances_students = attendances_students[config.ITEMS]

    return len(attendances_students) == 0


@blueprint.route('/students/attendances', methods=['GET'])
@requires_auth('/students/attendances')
def get_students():
    config_ori = deepcopy(app.config)
    # app.config['PAGINATION_DEFAULT'] = 9999
    app.config['DOMAIN']['attendances_tutors'].update({'embedded_fields': [
        'attendance',
        # 'attendance.module',
        # 'attendance.class',
        # 'attendance.class.branch',
    ]})
    lookup = {
        config.DATE_CREATED: {'$gte': (utc_now - timedelta(hours=12))},
        'tutor': app.auth.get_request_auth_value()
    }
    attendances, *_ = get_internal('attendances_tutors', **lookup)
    attendances = attendances[config.ITEMS]

    # app.config['DOMAIN']['classes_students'].update({'embedded_fields': [
    #     'student',
    # ]})
    for v in attendances:
        lookup = {
            'class': v['attendance']['class']
        }
        students, *_ = get_internal('classes_students', **lookup)
        students = students[config.ITEMS]

        # students = map(lambda v2: map_attendances(v, v2), students)
        students = filter(lambda v2: attendance_students(v2, v), students)
        students = list(students)
        v['students'] = students

    attendances = filter(lambda v: len(v['students']) > 0, attendances)
    attendances = {config.ITEMS: list(attendances)}

    app.config = config_ori
    return jsonify(attendances)


def _dormant_students(classes_students, student):
    for v in classes_students:
        if v['student'] == student[config.ID_FIELD]:
            return False
    return True


add_documentation({
    'paths': {'/students/dormant': {'get': {
        'summary': 'Retrieves one or more dormant users',
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
