from pytz import timezone
from datetime import timedelta, datetime, time
from pprint import pprint
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get, getitem
from eve.methods.put import put_internal
from eve.methods.post import post_internal
from eve.render import send_response
from eve_swagger import add_documentation
import humanize
from werkzeug.exceptions import NotFound

from shared.datetime import wib_tz, wib_now, utc_now, dow, dow_list

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def _last_attendance(class_):
    domain_ori = deepcopy(app.config['DOMAIN'])
    app.config['DOMAIN']['attendances_tutors'].update(
        {'embedded_fields': ['tutor']})

    utc_start_at = class_['startAtTs'].astimezone(timezone('UTC'))
    try:
        lookup = {
            '_created': {'$gte': utc_start_at},
            'class': class_['_id']
        }
        attendance, *_ = getitem('attendances', **lookup)
        lookup = {'attendance': attendance['_id']}
        attendances_tutors, *_ = get('attendances_tutors', **lookup)
    except NotFound:
        attendances_tutors = {'_items': []}

    app.config['DOMAIN'] = domain_ori
    return attendances_tutors


add_documentation({
    'paths': {'/schedules': {'put': {
        'summary': 'Replaces a Schedule document',
        'tags': ['Schedules'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/schedules', methods=['PUT'])
@requires_auth('/schedules')
def put_schedules():
    config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 9999
    app.config['DOMAIN']['classes'].update({'embedded_fields': [
        'branch',
        'module',
        'tutor',
    ]})
    classes, *_ = get('classes')
    classes = classes['_items']

    def parse_last_attendances(class_):
        class_.update({'last_attendances': _last_attendance(class_)})
        return class_
    classes = map(parse_last_attendances, classes)
    classes = {'_items': list(classes)}

    response = None
    resource = 'caches'
    payload = {
        'key': 'schedules',
        'value': classes
    }
    lookup = {
        'key': 'schedules'
    }
    try:
        response = put_internal(resource, payload, **lookup)
    except KeyError:
        response = post_internal(resource, payload)

    app.config = config_ori
    return jsonify({
        'response': response,
        'data': classes,
    })


add_documentation({
    'paths': {'/schedules': {'get': {
        'summary': 'Retrieves a Schedule document',
        'tags': ['Schedules'],
        "security": [{"JwtAuth": []}],
    }}}
})


def exclude_current_user_attendance(class_):

    if not class_['last_attendances'].get('_items'):
        return True

    items = class_['last_attendances']['_items']
    if len(items):
        for v2 in items:
            if v2['tutor'] == app.auth.get_request_auth_value():
                return False
    return True


def exclude_other_user_attendance(class_):
    if not class_['last_attendances'].get('_items'):
        return True

    items = class_['last_attendances']['_items']
    if len(items):
        if (class_['finishAtTs'] < utc_now):
            return False
    return True


date_range = [utc_now.date() + timedelta(days=v) for v in range(7)]


def _group(classes):
    group_list = []

    for date in date_range:
        group = {
            'date': date.isoformat(),
            'dateDay': date.day,
            'day': dow_list[date.weekday()],
            'text': '',
            '_items': [],
        }

        for class_ in classes:
            if date == class_['startAtTs'].date():
                text = humanize.naturaldelta(
                    timedelta(days=(class_['startAtTs'].date() - utc_now.date()).days))
                group['text'] = 'in %s' % text
                group['_items'].append(class_)

        if len(group['_items']):
            group_list.append(group)

    return group_list


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
    resource = 'caches'
    lookup = {'key': 'schedules'}
    row, *_ = getitem(resource, **lookup)
    classes = row['value']['_items']

    classes = filter(
        lambda v: v['finishAtTs'] > (wib_now - timedelta(hours=2)),
        classes)
    classes = filter(
        lambda v: v['finishAtTs'] < (wib_now + timedelta(days=5)),
        classes)
    classes = filter(exclude_current_user_attendance, classes)
    classes = filter(exclude_other_user_attendance, classes)
    classes = list(classes)

    classes.sort(key=lambda v: v['startAtTs'])
    classes = _group(classes)
    classes = {'_items': classes}

    return jsonify(classes)
