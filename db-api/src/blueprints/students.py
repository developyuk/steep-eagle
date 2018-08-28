from pytz import timezone
from pprint import pprint
from datetime import timedelta, datetime
from copy import deepcopy

from tables import SessionsTutors, SessionsStudents

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get, getitem

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 999
    app.config['DOMAIN']['sessions_tutors'].update({'embedded_fields': [
        'session',
        'session.class_',
        'session.class_.students',
        'session.class_.students.student',
        'session.class_.module',
        'session.class_.branch',
    ]})
    r = {
        '_created': '>=\'%s\'' % (datetime.now() - timedelta(hours=12)).strftime('%Y-%m-%d %H:%M:%S'),
        'tutor_id': app.auth.get_request_auth_value()
    }
    sessions, *_ = get('sessions_tutors', r)
    sessions = sessions['_items']

    def filter_session_students(v, st):
        # pprint(st)
        r = {
            '_created': '>=\'%s\'' % st['_created'].strftime('%Y-%m-%d %H:%M:%S'),
            'student_id': v['student']['id'],
            'session_id': st['session']['id'],
            'tutor_id': app.auth.get_request_auth_value()
        }
        sessions_students, *_ = get('sessions_students', r)
        sessions_students = sessions_students['_items']

        return len(sessions_students) == 0

    def filter_sessions(v):
        vlist = filter(lambda v2: filter_session_students(
            v2, v), v['session']['class_']['students'])
        v['session']['class_']['students'] = list(vlist)
        return v

    sessions = map(filter_sessions, sessions)
    sessions = list(sessions)
    sessions = filter(lambda v: len( v['session']['class_']['students']) > 0, sessions)
    sessions = list(sessions)

    # sessions = []
    app.config = app_config_ori
    return jsonify({
        '_items': sessions,
        '_meta': {
            'total': len(sessions)
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
