from pytz import timezone
from datetime import timedelta, datetime
from copy import deepcopy
# from pprint import pprint

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
import humanize
from eve.auth import requires_auth
from eve.methods import get, getitem
from eve.render import send_response
from . import wib_now, utc_now, onDay

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def last_attendance(class_):
    app.config['DOMAIN']['attendances'].update({'embedded_fields': [
        'attendance_tutors',
        'attendance_tutors.tutor',
    ]})

    utc_this = class_['start_at_ts'].astimezone(timezone('UTC'))
    lookup = {
        '_created': '>=\'%s\'' % utc_this.strftime('%Y-%m-%d'),
        'class_id': class_['id']
    }

    attendances, *_ = get('attendances', **lookup)
    attendances = attendances['_items']
    # attendances = attendances['_items'][0] if len(attendances['_items']) > 0 else []

    return attendances


def groupClass(classes):
    classes_group_list = []
    for class_ in classes:
        utc_this = class_['startAtTs'].astimezone(timezone('UTC'))
        date = utc_this.date().isoformat()
        ii = [i for i, j in enumerate(classes_group_list) if j['date'] == date]
        if not ii:
            classes_group_list.append({
                'date': date,
                'dateDay': class_['startAtTs'].day,
                'day': class_['day'],
                'text': 'in %s' % humanize.naturaldelta(timedelta(days=(utc_this.date() - utc_now.date()).days)),
                '_items': [class_],
            })
        else:
            classes_group_list[ii[0]]['_items'].append(class_)

    return classes_group_list


def exclude_current_user_attendance(v):
    if len(v['last_attendances']):
        for v2 in v['last_attendances'][0]['attendance_tutors']:
            if v2['tutor']['id'] == app.auth.get_request_auth_value():
                return False
    return True


def exclude_other_user_attendance(v):
    if len(v['last_attendances']):
        if (v['finish_at_ts'] < wib_now):
            return False
    return True


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
    resource = 'classes'
    response = get(resource)
    response = list(response)
    classes = response[0]['_items']

    classes = filter(
        lambda v: (v['finishAtTs'] + timedelta(hours=2)) > wib_now,
        classes)
    classes = filter(
        lambda v: v['finishAtTs'].date() < (
            wib_now + timedelta(days=5)).date(),
        classes)

    lookup = {'_id': app.auth.get_request_auth_value()}
    user, *_ = getitem('users', **lookup)
    def _exclude_dummies_non_tester(v):
        if 'tester' in user['username']:
            return 'dummies' in v['module']['name']
        else:
            return 'dummies' not in v['module']['name']
    classes = filter(_exclude_dummies_non_tester, classes)

    classes = list(classes)
    classes.sort(key=lambda v: v['startAtTs'])

    # def parse(v):
    #     v.update({'last_attendances': last_attendance(v)})
    #     return v
    # classes = map(parse, classes)
    # classes = filter(exclude_current_user_attendance, classes)
    # classes = filter(exclude_other_user_attendance, classes)
    classes = groupClass(classes)
    classes = list(classes)
    # classes = []
    response[0]['_items'] = classes

    return send_response(resource, response)


# @blueprint.after_request
# def add_header(response):
#     response.cache_control.max_age = app.config['CACHE_EXPIRES']
#     response.cache_control.public = True
#     response.cache_control.must_revalidate = True

#     now = datetime.now()
#     then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
#     response.headers['Expires'] = then
#     return response
