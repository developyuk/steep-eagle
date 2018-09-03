from datetime import timedelta, datetime
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get
from . import utc_now

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 999
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

    def filter_attendance_students(v, st):
        r = {
            '_created': '>=\'%s\'' % st['_created'].strftime('%Y-%m-%d'),
            'student_id': v['student']['id'],
            'attendance_id': st['attendance']['id'],
        }
        attendances_students, *_ = get('attendances_students', r)
        attendances_students = attendances_students['_items']

        return len(attendances_students) == 0

    def filter_attendances(v):
        vlist = filter(lambda v2: filter_attendance_students(
            v2, v), v['attendance']['class_']['students'])
        v['attendance']['class_']['students'] = list(vlist)
        return v

    attendances = map(filter_attendances, attendances)
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


@blueprint.after_request
def add_header(response):
    response.cache_control.max_age = app.config['CACHE_EXPIRES']
    response.cache_control.public = True
    response.cache_control.must_revalidate = True

    now = datetime.now()
    then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
    response.headers['Expires'] = then
    return response
