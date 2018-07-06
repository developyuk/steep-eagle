from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from tables import ClassesTs, Sessions
from datetime import timedelta, datetime
import humanize
from eve.auth import requires_auth
from pytz import timezone

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
  # .filter(ClassesTs.finish_at_ts < (datetime.now(timezone('UTC')) + timedelta(days=5))) \
  classes = app.data.driver.session.query(ClassesTs) \
    .all()

  # data = map(lambda x: x.to_dict(), classes)

  classes_list = []
  for class_ in classes:
    w = class_.to_dict()

    sessions = app.data.driver.session.query(Sessions) \
      .filter(Sessions._created >= class_.start_at_ts) \
      .filter(Sessions.class_id == class_.id) \
      .all()

    sessions_list = []
    for session in sessions:
      ww = session.to_dict()
      session_tutors_list = []
      for vvv in session.session_tutors:
        www = vvv.to_dict()
        www.update({
          'tutor': vvv.tutor.to_dict(),
        })
        www['tutor'].update({
          'profile': vvv.tutor.profile.to_dict(),
        })
        session_tutors_list.append(www)

      ww.update({
        'session_tutors': session_tutors_list
      })
      sessions_list.append(ww)

    w.update({
      'branch': class_.branch.to_dict(),
      'program_module': class_.program_module.to_dict(),
      'q': class_.q,
      'tutor': class_.tutor.to_dict(),
      'finish_at_ts': class_.finish_at_ts,
      'start_at_ts': class_.start_at_ts,
      'last_sessions': sessions_list,
    })
    w['tutor'].update({
      'profile': class_.tutor.profile.to_dict(),
    })
    w['program_module'].update({
      'module': class_.program_module.module.to_dict(),
    })

    classes_list.append(w)

  classes_list.sort(key=lambda v: v['start_at_ts'])
  classes_list = list(filter(lambda v: v['finish_at_ts'] < (datetime.now(timezone('UTC')) + timedelta(days=5)), classes_list))
  classes_group_list = []

  for class_ in classes_list:
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

  # classes_group_list = []
  return jsonify({
    '_items': classes_group_list,
    'meta': {'total': len(classes_group_list)}
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
