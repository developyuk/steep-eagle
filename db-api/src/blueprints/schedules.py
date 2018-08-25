from pytz import timezone
from pprint import pprint
from datetime import timedelta, datetime
import json

from tables import Users, ClassesTs, Sessions

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
import humanize
from eve.auth import requires_auth
from eve.methods import get, getitem

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def last_session(class_):
    r = {
        '_created': '>=\'%s\'' % class_['start_at_ts'].strftime('%Y-%m-%d'),
        'class_id': class_['id']
    }
    # pprint(r)
    sessions, *_ = get('sessions', r)
    sessions = sessions['_items']
    # pprint(sessions)
    # sessions = app.data.driver.session.query(Sessions).filter(Sessions._created >= class_['start_at_ts']).filter(Sessions.class_id == class_['id']).all()

    return sessions


def groupClass(classes):
    classes_group_list = []
    for class_ in classes:
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
    classes, *_ = get('classes_ts')
    # document = app.data.driver.session.query(ClassesTs).all()
    # classes = list(classes)
    # pprint(classes)
    classes = classes['_items']

    user, *_ = getitem('users', {'id': app.auth.get_request_auth_value()})
    # user = app.data.driver.session.query(Users).get(app.auth.get_request_auth_value())

    def exclude_dummies_non_tester(v):
        if 'tester' not in user['username']:
            return 'dummies' not in v['module']['name']
        else:
            return True

    classes = filter(lambda v: v['finish_at_ts'] > (datetime.now(
        timezone('Asia/Jakarta')) - timedelta(hours=2)), classes)
    classes = filter(lambda v: v['finish_at_ts'].date() < (datetime.now(
        timezone('Asia/Jakarta')) + timedelta(days=5)).date(), classes)
    classes = filter(exclude_dummies_non_tester, classes)

    classes = list(classes)
    classes.sort(key=lambda v: v['start_at_ts'])

    def parse(v):
        v.update({'last_sessions': last_session(v)})
        return v

    classes = [parse(v) for v in classes]
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
