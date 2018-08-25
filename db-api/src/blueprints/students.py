from pytz import timezone
from pprint import pprint
from datetime import timedelta, datetime

from tables import SessionsTutors, SessionsTutorsStudents

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get, getitem

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
    # sessions = app.data.driver.session.query(SessionsTutors) \
    #   .filter(SessionsTutors._created >= datetime.now() - timedelta(hours=12)) \
    #   .filter(SessionsTutors.tutor_id == app.auth.get_request_auth_value()) \
    #   .all()
    r = {
        '_created': '>=\'%s\'' % (datetime.now() - timedelta(hours=12)).strftime('%Y-%m-%d %H:%M:%S'),
        'tutor_id': app.auth.get_request_auth_value()
    }
    sessions, *_ = get('sessions_tutors', r)
    pprint(sessions)
    sessions = sessions['_items']

    def filter_session_students(v, st):
        # sessions_students = app.data.driver.session.query(SessionsTutorsStudents) \
        #     .filter(SessionsTutorsStudents.student_id == v.student.id) \
        #     .filter(SessionsTutorsStudents._created >= st._created) \
        #     .all()

        r = {
            '_created': '>=\'%s\'' % st['_created'].strftime('%Y-%m-%d %H:%M:%S'),
            'student_id': v['student']
        }
        sessions_students, *_ = get('sessions_tutors_students', r)
        sessions_students = sessions_students['_items']

        return not len(sessions_students)

    def filter_sessions(v):
        pprint(v['session']['class_'])
        vlist = filter(lambda v2: filter_session_students(v2, v), v['session']['class_']['students'])
        vlist = list(vlist)
        return len(vlist)

    sessions = filter(filter_sessions, sessions)
    sessions = list(sessions)

    # def parse(v):
    #     w = dict(v).copy()
    #     # w.update({
    #     #     'session': dict(v.session)
    #     # })
    #     # w['session'].update({
    #     #     'class': dict(v.session.class_)
    #     # })
    #     # w['session']['class'].update({
    #     #     'branch': dict(v.session.class_.branch),
    #     #     'module': dict(v.session.class_.module),
    #     #     # 'students': map(lambda vv: dict(vv), v.session.class_.students),
    #     # })

    #     # for ii, vv in enumerate(v.session.class_.students):
    #     #     w['session']['class']['students'][ii].update({
    #     #         'student': dict(vv.student)
    #     #     })
    #     return w

    # sessions = list(map(parse, sessions))

    # sessions = []
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
