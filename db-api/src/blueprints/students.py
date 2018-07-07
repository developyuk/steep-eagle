from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from eve.auth import requires_auth
from datetime import timedelta, datetime
from tables import SessionsTutors, Sessions, ClassesTs
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

  def merge(v):
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
      w['session']['class']['students'][ii]['student'].update({
        'profile': dict(vv.student.profile)
      })
    return w

  sessions = map(merge, sessions)

  # sessions = []
  return jsonify({
    '_items': sessions,
    '_meta': {
      'total': len(sessions)
    }
  })

# @blueprint.after_request
# def add_header(response):
#   response.cache_control.max_age = app.config['CACHE_EXPIRES']
#   response.cache_control.public = True
#   response.cache_control.must_revalidate = True
#
#   now = datetime.now()
#   then = now + timedelta(days=app.config['CACHE_EXPIRES'])
#   response.headers['Expires'] = then
#   return response
