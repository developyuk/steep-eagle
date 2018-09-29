from datetime import timedelta
from copy import deepcopy
from pprint import pprint

from flask_cors import CORS
from eve_swagger import add_documentation
from flask import current_app as app, jsonify, Blueprint
from eve.auth import requires_auth
from eve.methods import get, getitem
from eve.methods.put import put_internal
from eve.methods.post import post_internal
from eve.render import send_response
from werkzeug.exceptions import NotFound

from shared.datetime import utc_now

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def attendance_students(item, at):
    lookup = {
        '_created': {'$gte': at['_created'].replace(hour=0, minute=0, second=0, microsecond=0)},
        'student': item['student']['_id'],
        'attendance': at['attendance']['_id'],
    }
    attendances_students, *_ = get('attendances_students', **lookup)
    attendances_students = attendances_students['_items']

    return len(attendances_students) == 0


@blueprint.route('/students/attendances', methods=['PUT'])
@requires_auth('/students/attendances')
def put_students():
    config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 9999
    app.config['DOMAIN']['attendances_tutors'].update({'embedded_fields': [
        'attendance',
        'attendance.module',
        'attendance.class',
        'attendance.class.branch',
    ]})
    lookup = {
        '_created': {'$gte': (utc_now - timedelta(hours=12))},
        'tutor': app.auth.get_request_auth_value()
    }
    attendances, *_ = get('attendances_tutors', **lookup)
    attendances = attendances['_items']

    app.config['DOMAIN']['classes_students'].update({'embedded_fields': [
        'student',
    ]})
    for v in attendances:
        lookup = {
            'class': v['attendance']['class']['_id']
        }
        students, *_ = get('classes_students', **lookup)
        students = students['_items']

        # students = map(lambda v2: map_attendances(v, v2), students)
        students = filter(lambda v2: attendance_students(v2, v), students)
        students = list(students)
        v['students'] = students

    attendances = filter(lambda v: len(v['students']) > 0, attendances)
    attendances = {'_items': list(attendances)}

    response = None
    resource = 'caches'
    payload = {
        'key': 'students_attendances_%s' % app.auth.get_request_auth_value(),
        'value': attendances
    }
    lookup = {
        'key': 'students_attendances_%s' % app.auth.get_request_auth_value()
    }
    try:
        response = put_internal(resource, payload, **lookup)
    except KeyError:
        response = post_internal(resource, payload)

    app.config = config_ori
    return jsonify({
        'response': response,
        'data': attendances,
    })


@blueprint.route('/students/attendances', methods=['GET'])
@requires_auth('/students/attendances')
def get_students():
    resource = 'caches'
    lookup = {'key': 'students_attendances_%s' % app.auth.get_request_auth_value()}
    try:
        row, *_ = getitem(resource, **lookup)
        attendances = row['value']
    except NotFound:
        attendances = {'_items': []}

    return jsonify(attendances)


def _dormant_students(v):
    try:
        _ = getitem('classes_students', **{'student': v['_id']})
        return False
    except NotFound:
        return True


add_documentation({
    'paths': {'/students/dormant': {'get': {
        'summary': 'Retrieves one or more dormant students',
        "responses": {
            "200": {
                "description": "Student document fetched successfully",
                "schema": {"$ref": "#/definitions/Student"}
            }
        },
        'tags': ['Student'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/students/dormant', methods=['GET'])
@requires_auth('/students/dormant')
def students_dormant():
    resource = 'students'
    response = get(resource)
    response = list(response)
    students = response[0]['_items']

    students = filter(_dormant_students, students)
    response[0]['_items'] = list(students)

    return send_response(resource, response)
