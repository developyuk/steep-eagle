from pytz import timezone
from datetime import timedelta, datetime, time
from pprint import pprint
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint, request
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get, getitem
# from eve.methods.put import put_internal
# from eve.methods.post import post_internal
# from eve.render import send_response
from eve_swagger import add_documentation
from eve.utils import config
import humanize
from werkzeug.exceptions import NotFound

from shared.datetime import wib_tz, wib_now, utc_now, dow, dow_list

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def _last_attendance(class_):
    # domain_ori = deepcopy(app.config['DOMAIN'])
    # app.config['DOMAIN']['attendances_tutors'].update(
    #     {'embedded_fields': ['tutor']})

    utc_start_at = class_['start'].astimezone(timezone('UTC'))
    try:
        lookup = {
            '_created': {'$gte': utc_start_at},
            'class': class_[config.ID_FIELD]
        }
        attendance, *_ = getitem('attendances', **lookup)
        lookup = {'attendance': attendance[config.ID_FIELD]}
        attendances_tutors, *_ = get('attendances_tutors', **lookup)
    except NotFound:
        attendances_tutors = {config.ITEMS: []}

    # app.config['DOMAIN'] = domain_ori
    return attendances_tutors


add_documentation({
    'paths': {'/schedules': {'get': {
        'summary': 'Retrieves a Schedule document',
        'tags': ['Classe'],
        "security": [{"JwtAuth": []}],
    }}}
})


def exclude_current_user_attendance(class_):
    if not class_['last_attendances'].get(config.ITEMS):
        return True

    items = class_['last_attendances'][config.ITEMS]
    if len(items):
        for v2 in items:
            if v2['tutor'] == app.auth.get_request_auth_value():
                return False
    return True


def exclude_other_user_attendance(class_):
    if not class_['last_attendances'].get(config.ITEMS):
        return True

    items = class_['last_attendances'][config.ITEMS]
    if len(items):
        if (class_['finishAtTs'] < utc_now):
            return False
    return True


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
    _page = int(request.args.get("_page", 1))
    _max_results = int(request.args.get("_max_results", 1))

    date_range = ((wib_now.date()+timedelta(days=(_page-1)*_max_results)) + timedelta(days=v)
                  for v in range(_max_results))

    classes, *_ = get('classes')
    classes = classes[config.ITEMS]
    # pprint(classes)

    data = []
    for date in date_range:
        d = {
            'date': date.isoformat(),
            'delta': humanize.naturaldelta(
                timedelta(days=(date - wib_now.date()).days)),
            'day': date.day,
            'weekday': dow_list[date.weekday()],
            config.ITEMS: []
        }
        for class_ in classes:
            recurrence = class_['schedule']['recurrence']
            class_['start'] = class_['start'].replace(
                day=date.day, month=date.month, year=date.year)
            class_['finish'] = class_['finish'].replace(
                day=date.day, month=date.month, year=date.year)

            if recurrence['freq'] == 'daily':
                pass
            if recurrence['freq'] == 'weekly':
                if recurrence['interval'] == 1:
                    if dow_list[date.weekday()] in recurrence['byday']:
                        d[config.ITEMS].append(class_)

        if len(d[config.ITEMS]) > 0:
            d[config.ITEMS].sort(key=lambda v: v['start'])

            for v in d[config.ITEMS]:
                v['last_attendances'] = _last_attendance(v)

            d[config.ITEMS] = filter(
                exclude_current_user_attendance, d[config.ITEMS])
            d[config.ITEMS] = filter(
                exclude_other_user_attendance, d[config.ITEMS])
            d[config.ITEMS] = list(d[config.ITEMS])
            data.append(d)

    return jsonify({
        config.ITEMS: data,
        config.META: {
            'page': _page,
            'max_results': _max_results,
        }
    })
