from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from datetime import timedelta, datetime
from tables import SessionsTutors, SessionsTutorsStudents
from pytz import timezone
from pprint import pprint

blueprint = Blueprint('students', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
  sessions = app.data.driver.session.query(SessionsTutors) \
    .filter(SessionsTutors._created >= datetime.now() - timedelta(hours=12)) \
    .filter(SessionsTutors.tutor_id == app.auth.get_request_auth_value()) \
    .all()

  def filter_session_students(v, st):
    sessions_students = app.data.driver.session.query(SessionsTutorsStudents) \
      .filter(SessionsTutorsStudents.student_id == v.student.id) \
      .filter(SessionsTutorsStudents._created >= st._created) \
      .all()

    return not len(sessions_students)

  def filter_sessions(v):
    v.session.class_.students = filter(lambda v2: filter_session_students(v2, v), v.session.class_.students)
    return len(v.session.class_.students)

  sessions = filter(filter_sessions, sessions)

  def parse(v):
    w = dict(v).copy()
    w.update({
      'session': dict(v.session)
    })
    w['session'].update({
      'class': dict(v.session.class_)
    })
    w['session']['class'].update({
      'branch': dict(v.session.class_.branch),
      'program_module': dict(v.session.class_.program_module),
      'students': map(lambda vv: dict(vv), v.session.class_.students),
    })
    w['session']['class']['program_module'].update({
      'module': dict(v.session.class_.program_module.module)
    })

    for ii, vv in enumerate(v.session.class_.students):
      w['session']['class']['students'][ii].update({
        'student': dict(vv.student)
      })
    return w

  sessions = map(parse, sessions)

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
