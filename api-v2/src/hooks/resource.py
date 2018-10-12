from pprint import pprint

from flask import current_app as app, abort
from eve.methods.get import getitem_internal
from eve.methods.post import post_internal
from eve.utils import config
from shared.datetime import wib_now, wib_tz, dow_list
from werkzeug.exceptions import NotFound


def gen_students_total(item):
    item['studentsTotal'] = 0


def conv_fetch_schedule(item, type_='resource'):
    item['_start'] = item['start'].astimezone(wib_tz).strftime('%H:%M')
    item['_finish'] = item['finish'].astimezone(wib_tz).strftime('%H:%M')
    item['_day'] = dow_list[item['start'].astimezone(wib_tz).weekday()]
    recurrence = item['schedule']['recurrence']
    if type_ == 'resource':
        item['_schedule'] = '%s on %s' % (
            recurrence['freq'], recurrence['byday'][0])
    else:
        item['_schedule'] = recurrence['byday'][0]


def on_fetched_resource(resource_name, response):
    if resource_name == 'classes':
        for v in response[config.ITEMS]:
            gen_students_total(v)
            conv_fetch_schedule(v, 'resource')


def on_fetched_item(resource_name, response):
    if resource_name == 'classes':
        gen_students_total(response)
        conv_fetch_schedule(response, 'item')


def conv_put_recurr(item):
    start_at = item['startAt'].split(':')
    start_at = wib_now.replace(
        hour=int(start_at[0]), minute=int(start_at[1]), second=0, microsecond=0)
    finish_at = item['finishAt'].split(':')
    finish_at = wib_now.replace(
        hour=int(finish_at[0]), minute=int(finish_at[1]), second=0, microsecond=0)

    return {
        'start': start_at,
        'finish': finish_at,
        'schedule': {
            'recurrence': {
                'interval': 1,
                'freq': 'weekly',
                'byday': [item['day']],
                'until': None,
                'count': None,
            },
            'include': None,
            'exclude': None,
        }
    }


def on_update(resource_name, updates, original):
    if resource_name == 'classes':
        d = conv_put_recurr(updates)
        updates.update(d)
        del updates['day']
        del updates['startAt']
        del updates['finishAt']


def on_inserted(resource_name, items):
    pass


def on_insert(resource_name, items):
    if resource_name == 'classes':
        for i, v in enumerate(items):
            d = conv_put_recurr(v)
            items[i].update(d)
            del items[i]['day']
            del items[i]['startAt']
            del items[i]['finishAt']

    if resource_name == 'attendances_tutors':
        for i, v in enumerate(items):
            class_ = v.pop('class', None)

            lookup = {config.ID_FIELD: class_}
            class_, *_ = getitem_internal('classes', **lookup)
            date = wib_now.date()

            # pprint('class_')
            # pprint(class_)

            try:
                lookup = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module'],
                    config.DATE_CREATED: {'$gte': class_['start'].astimezone(wib_tz).replace(
                        day=date.day, month=date.month, year=date.year)},
                }
                attendance, *_ = getitem_internal('attendances', **lookup)
            except NotFound:
                payload = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module']
                }
                attendance, *_ = post_internal('attendances', payload)

            # pprint('attendance')
            # pprint(attendance)

            try:
                lookup = {
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                    config.DATE_CREATED: {'$gte': attendance[config.DATE_CREATED]},
                }
                *_, status = getitem_internal('attendances_tutors', **lookup)

                if status == 200:
                    abort(422, description='exist')
            except NotFound:
                # pprint(items)
                items[i].update({
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                })
        # pprint(items)
