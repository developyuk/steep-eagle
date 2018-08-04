from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from tables import Users, ClassesTs, Sessions
from datetime import timedelta, datetime
import humanize
from eve.auth import requires_auth
from pytz import timezone
from pprint import pprint

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def last_session(class_):
  sessions = app.data.driver.session.query(Sessions) \
    .filter(Sessions._created >= class_.start_at_ts) \
    .filter(Sessions.class_id == class_.id) \
    .all()

  def parseTutor(v):
    w = dict(v)
    w.update({
      'tutor': dict(v.tutor),
    })
    return w

  def parse(v):
    w = dict(v)

    w.update({
      'session_tutors': map(parseTutor, v.session_tutors)
    })
    return w

  sessions = map(parse, sessions)

  return sessions


def groupClass(classes):
  classes_group_list = []
  for class_ in classes:
    # pass
    date = class_['start_at_ts'].date().isoformat()
    ii = [i for i, j in enumerate(classes_group_list) if j['date'] == date]
    if not ii:
      classes_group_list.append({
        'date': date,
        'dateDay': class_['start_at_ts'].day,
        'day': class_['day'],
        'text': 'in %s' % humanize.naturaldelta(
          timedelta(days=(class_['start_at_ts'] - datetime.now(timezone('UTC'))).days)),
        '_items': [class_],
      })
    else:
      classes_group_list[ii[0]]['_items'].append(class_)

  return classes_group_list


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
  classes = app.data.driver.session.query(ClassesTs) \
    .all()

  pprint(app.auth.get_request_auth_value())
  user = app.data.driver.session.query(Users).get(app.auth.get_request_auth_value())
  pprint(user)

  def exclude_dummies_non_tester(v):
    if 'tester' not in user.username:
      return 'dummies' not in v.module.name
    else:
      return True

  classes = filter(lambda v: v.finish_at_ts > (datetime.now(timezone('UTC')) - timedelta(hours=2)), classes)
  classes = filter(lambda v: v.finish_at_ts.date() < (datetime.now(timezone('UTC')) + timedelta(days=5)).date(),
                   classes)
  classes = filter(exclude_dummies_non_tester, classes)
  classes.sort(key=lambda v: v.start_at_ts)

  def parse(v):
    w = dict(v).copy()

    w.update({
      'branch': dict(v.branch),
      'module': dict(v.module),
      'tutor': dict(v.tutor),
      'finish_at_ts': v.finish_at_ts,
      'start_at_ts': v.start_at_ts,
      'last_sessions': last_session(v),
    })
    return w

  classes = map(parse, classes)
  classes = groupClass(classes)
  # classes = []
  return jsonify({
    '_items': classes,
    'meta': {'total': len(classes)}
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
