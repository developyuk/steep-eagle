
from datetime import timedelta
from pprint import pprint

from flask import current_app as app, jsonify, Blueprint, request
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods.get import get_internal
# from eve.methods.put import put_internal
# from eve.methods.post import post_internal
# from eve.render import send_response
from eve_swagger import add_documentation
from eve.utils import config
import humanize
from pytz import utc

from shared.datetime import wib_tz, wib_now, utc_now, dow_list, after_request_cache

blueprint = Blueprint('schedules', __name__)
CORS(blueprint, max_age=timedelta(seconds=10))


def _last_attendance(attendances, attendances_tutors, class_):
    # domain_ori = deepcopy(app.config['DOMAIN'])
    # app.config['DOMAIN']['attendances_tutors'].update(
    #     {'embedded_fields': ['tutor']})

    # utc_start_at = class_['start'].astimezone(timezone('UTC'))
    # try:
    #     lookup = {
    #         config.DATE_CREATED: {'$gte': utc_start_at},
    #         'class': class_[config.ID_FIELD]
    #     }
    #     attendance, *_ = getitem('attendances', **lookup)
    #     lookup = {'attendance': attendance[config.ID_FIELD]}
    #     attendances_tutors, *_ = get('attendances_tutors', **lookup)
    # except NotFound:
    #     attendances_tutors = {config.ITEMS: []}

    # app.config['DOMAIN'] = domain_ori

    data = {config.ITEMS: []}

    if not len(attendances):
        return data

    if not len(attendances_tutors):
        return data

    utc_start_at = class_['start'].astimezone(utc)

    class_attendance = {}
    for v in attendances:
        if v['class'] == class_[config.ID_FIELD] and v[config.DATE_CREATED].astimezone(utc) >= utc_start_at:
            class_attendance = v
            break

    if not class_attendance.get(config.ID_FIELD):
        return data

    class_attendances_tutors = []
    for v in attendances_tutors:
        if v['attendance'] == class_attendance[config.ID_FIELD]:
            class_attendances_tutors.append(v)

    data = {config.ITEMS: class_attendances_tutors}

    return data


add_documentation({
    'paths': {'/schedules': {'get': {
        'summary': 'Retrieves a Schedule document',
        'tags': ['Classe'],
        "security": [{"JwtAuth": []}],
    }}}
})


def exclude_current_user_attendance(class_):
    if not class_.get('last_attendances'):
        return True

    if not class_['last_attendances'].get(config.ITEMS):
        return True

    items = class_['last_attendances'][config.ITEMS]
    if len(items):
        for v2 in items:
            if v2['tutor'] == app.auth.get_request_auth_value():
                return False
    return True


def exclude_other_user_attendance(class_):
    if not class_.get('last_attendances'):
        return True

    if not class_['last_attendances'].get(config.ITEMS):
        return True

    items = class_['last_attendances'][config.ITEMS]
    if len(items):
        if (class_['finish'] < utc_now):
            return False
    return True


def conv_time(class_, date):
    class_['start'] = class_['start'].astimezone(wib_tz).replace(
        day=date.day, month=date.month, year=date.year)
    class_['finish'] = class_['finish'].astimezone(wib_tz).replace(
        day=date.day, month=date.month, year=date.year)


def embed_tutor(tutors, item):
    if not len(item['last_attendances'][config.ITEMS]):
        return item

    for v in item['last_attendances'][config.ITEMS]:
        for v2 in tutors:
            if v2[config.ID_FIELD] == v['tutor']:
                v['tutor'] = v2
                break

    return item


@blueprint.route('/schedules', methods=['GET'])
@requires_auth('/schedules')
def schedules():
    _page = int(request.args.get("_page", 1))
    _max_results = int(request.args.get("_max_results", 1))

    date_range = ((wib_now.date()+timedelta(days=(_page-1)*_max_results)) + timedelta(days=v)
                  for v in range(_max_results))

    classes, *_ = get_internal('classes')
    classes = classes[config.ITEMS]

    attendances, *_ = get_internal('attendances')
    attendances = attendances[config.ITEMS]

    attendances_tutors, *_ = get_internal('attendances_tutors')
    attendances_tutors = attendances_tutors[config.ITEMS]

    lookup = {
        'role': 'tutor'
    }
    tutors, *_ = get_internal('users', **lookup)
    tutors = tutors[config.ITEMS]

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

            if recurrence['freq'] == 'daily':
                pass

            if recurrence['freq'] == 'weekly':
                if recurrence['interval'] == 1:
                    if dow_list[date.weekday()] in recurrence['byday']:
                        conv_time(class_, date)
                        d[config.ITEMS].append(class_)

        if len(d[config.ITEMS]) > 0:
            d[config.ITEMS].sort(key=lambda v: v['start'])

            for v in d[config.ITEMS]:
                v['last_attendances'] = _last_attendance(
                    attendances, attendances_tutors, v)

            d[config.ITEMS] = filter(
                exclude_current_user_attendance, d[config.ITEMS])
            d[config.ITEMS] = filter(
                exclude_other_user_attendance, d[config.ITEMS])
            d[config.ITEMS] = map(
                lambda v: embed_tutor(tutors, v), d[config.ITEMS])
            d[config.ITEMS] = list(d[config.ITEMS])
            data.append(d)

    return jsonify({
        config.ITEMS: data,
        config.META: {
            'page': _page,
            'max_results': _max_results,
        }
    })


@blueprint.after_request
def add_header(response):
    return after_request_cache(response)
