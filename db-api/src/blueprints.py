from flask import current_app as app, jsonify, Blueprint, request
from tables import Users, ClassesTs
from pprint import pprint
import humanize
import datetime

blueprint = Blueprint('prefix_uri', __name__)


@blueprint.route('/classes_ts_group_by_dow', methods=['GET'])
def classes_ts_group_by_dow():
  classes = app.data.driver.session.query(ClassesTs).all()
  # data = map(lambda x: x.to_dict(), classes)

  classes_list = []
  for v in classes:
    w = v.to_dict()
    w.update({
      'branch': v.branch.to_dict(),
      'tutor': v.tutor.to_dict(),
      'program_module': v.program_module.to_dict(),
      'q': v.q,
      'finish_at_ts': v.finish_at_ts,
      'start_at_ts': v.start_at_ts,
    })
    w['tutor'].update({
      'profile': v.tutor.profile.to_dict(),
    })
    w['program_module'].update({
      'module': v.program_module.module.to_dict(),
      'program': v.program_module.program.to_dict(),
    })
    w['program_module']['program'].update({
      'type': v.program_module.program.type.to_dict(),
    })
    classes_list.append(w)

  classes_list.sort(key=lambda x: x['start_at_ts'])
  classes_group_list = []

  for v in classes_list:
    # pass
    date = v['start_at_ts'].date().isoformat()
    ii = [i for i, j in enumerate(classes_group_list) if j['date'] == date]
    if not ii:
      classes_group_list.append({
        'date': date,
        'day': v['day'],
        'text': 'in %s' % humanize.naturaldelta(datetime.timedelta(days=(v['start_at_ts'].replace(tzinfo=None) - datetime.datetime.now()).days)),
        '_items': [v],
      })
    else:
      classes_group_list[ii[0]]['_items'].append(v)

  return jsonify({
    '_items': classes_group_list,
    'meta': {'total': len(classes_group_list)}
  })


@blueprint.route('/sign', methods=['POST'])
def sign():
  data = request.values or request.get_json()

  try:
    username = data.get('username')
    user = app.data.driver.session.query(Users).filter(Users.username == username).first()
    return jsonify({'token': user.sign()})
  except AttributeError:
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': 'Please provide proper credentials',
        "code": 400
      }
    }), 400


@blueprint.route('/auth', methods=['GET'])
def auth():
  data = request.headers
  try:
    token = data.get('Authorization')
    token = token.split(' ')[1]
    user = Users.auth(token)
    return jsonify(user)
  except (AttributeError, TypeError):
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': 'Please provide proper credentials',
        "code": 400
      }
    }), 400
