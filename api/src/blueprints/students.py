from datetime import timedelta, datetime
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get, getitem
from . import utc_now
from werkzeug.exceptions import NotFound

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def filter_attendance_students(v, st):
    r = {
        '_created': '>=\'%s\'' % st['_created'].strftime('%Y-%m-%d'),
        'student_id': v['student']['id'],
        'attendance_id': st['attendance']['id'],
    }
    attendances_students, *_ = get('attendances_students', r)
    attendances_students = attendances_students['_items']

    return len(attendances_students) == 0


def filter_students_is_deleted(v):
    r = {
        'id': v['student']['id'],
        'role': 'student',
        'is_deleted': False,
    }
    try:
        student, *_ = getitem('users', r)
        return True
    except NotFound as e:
        return False


def map_attendances(v):
    vlist = filter(lambda v2: filter_attendance_students(
        v2, v), v['attendance']['class_']['students'])
    vlist = filter(filter_students_is_deleted,
                   v['attendance']['class_']['students'])
    vlist = list(vlist)
    v['attendance']['class_']['students'] = vlist
    return v


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 999
    # app.config['DOMAIN']['attendances_students'].update({'embedded_fields': [
    #     'student'
    # ]})
    app.config['DOMAIN']['attendances_tutors'].update({'embedded_fields': [
        'attendance',
        'attendance.class_',
        'attendance.class_.students',
        'attendance.class_.students.student',
        'attendance.class_.module',
        'attendance.class_.branch',
    ]})
    r = {
        '_created': '>=\'%s\'' % (utc_now - timedelta(hours=12)).strftime('%Y-%m-%d %H:%M:%S'),
        'tutor_id': app.auth.get_request_auth_value()
    }
    attendances, *_ = get('attendances_tutors', r)
    attendances = attendances['_items']

    attendances = map(map_attendances, attendances)
    attendances = list(attendances)
    attendances = filter(lambda v: len(
        v['attendance']['class_']['students']) > 0, attendances)
    attendances = list(attendances)

    # attendances = []
    app.config = app_config_ori
    return jsonify({
        '_items': attendances,
        '_meta': {
            'total': len(attendances)
        }
    })


def dormant_students(v):
    try:
        student, *_ = getitem('class_students', **{'student_id': v['id']})
        return False
    except NotFound as e:
        return True


@blueprint.route('/students/dormant', methods=['GET'])
@requires_auth('/students/dormant')
def students_dormant():
    # app_config_ori = deepcopy(app.config)
    # app.config['PAGINATION_DEFAULT'] = 9999
    students, *_ = get('users', **{'role': 'student', 'is_deleted': False})
    students = students['_items']

    students = filter(dormant_students, students)
    students = list(students)

    # app.config = app_config_ori
    return jsonify({
        '_items': students,
        '_meta': {
            'total': len(students)
        }
    })


@blueprint.after_request
def add_header(response):
    response.cache_control.max_age = app.config['CACHE_EXPIRES']
    response.cache_control.public = True
    response.cache_control.must_revalidate = True

    now = datetime.now()
    then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
    response.headers['Expires'] = then
    return response
