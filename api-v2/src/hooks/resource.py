from pprint import pprint

from pytz import utc
from datetime import datetime
from flask import current_app as app, abort
from eve.methods import getitem, post, delete
from eve.methods.get import getitem_internal
from eve.methods.post import post_internal
from eve.utils import config
from shared.datetime import utc_now, on_day, wib_now, wib_tz
from werkzeug.exceptions import NotFound

import blueprints


def gen_students_total(item):
    item['studentsTotal'] = 0


def conv_start_finish(item, type_='resource'):
    item['_start'] = item['start'].astimezone(wib_tz).strftime('%H:%M')
    item['_finish'] = item['finish'].astimezone(wib_tz).strftime('%H:%M')
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
            conv_start_finish(v, 'resource')


def on_fetched_item(resource_name, response):
    if resource_name == 'classes':
        gen_students_total(response)
        conv_start_finish(response, 'item')


def on_inserted(resource_name, items):
    pass


def on_insert(resource_name, items):
    if resource_name == 'classes':
        for i, v in enumerate(items):
            start_at = v['startAt'].split(':')
            start_at = wib_now.replace(
                hour=int(start_at[0]), minute=int(start_at[1]), second=0, microsecond=0)
            finish_at = v['finishAt'].split(':')
            finish_at = wib_now.replace(
                hour=int(finish_at[0]), minute=int(finish_at[1]), second=0, microsecond=0)

            items[i].update({
                'start': start_at,
                'finish': finish_at,
                'schedule': {
                    'recurrence': {
                        'interval': 1,
                        'freq': 'weekly',
                        'byday': [v['day']],
                        'until': None,
                        'count': None,
                    },
                    'include': None,
                    'exclude': None,
                }
            })
            del items[i]['day']
            del items[i]['startAt']
            del items[i]['finishAt']

    if resource_name == 'attendances_tutors':
        for i, v in enumerate(items):
            class_ = v.pop('class', None)

            lookup = {config.ID_FIELD: class_}
            class_, *_ = getitem_internal('classes', **lookup)

            try:
                lookup = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module'],
                    '_created': {'$gte': class_['startAtTs']},
                }
                attendance, *_ = getitem('attendances', **lookup)
            except NotFound:
                payload = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module']
                }
                attendance, *_ = post_internal('attendances', payload)

            try:
                lookup = {
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                    '_created': {'$gte': attendance['_created']},
                }
                *_, status = getitem('attendances_tutors', **lookup)

                if status == 200:
                    abort(422, description='exist')
            except NotFound:
                items[i].update({
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                })
