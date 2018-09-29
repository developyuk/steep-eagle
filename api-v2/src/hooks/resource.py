from pprint import pprint

from pytz import timezone
from datetime import datetime
from flask import current_app as app, abort
from eve.methods import getitem, post
from eve.methods.get import getitem_internal
from eve.methods.post import post_internal
from shared.datetime import utc_now, on_day
from werkzeug.exceptions import NotFound

import blueprints


def _gen_students_total(item):
    item['studentsTotal'] = 0


def _gen_timestamp(item):

    dt = datetime.strptime('%sT%s' % (
        on_day(item['day']).date(), item['startAt']), '%Y-%m-%dT%H:%M')
    item['startAtTs'] = timezone('Asia/Jakarta').localize(dt)

    dt = datetime.strptime('%sT%s' % (
        on_day(item['day']).date(), item['finishAt']), '%Y-%m-%dT%H:%M')
    item['finishAtTs'] = timezone('Asia/Jakarta').localize(dt)


def on_fetched_resource(resource_name, response):
    if resource_name == 'classes':
        for v in response['_items']:
            _gen_timestamp(v)
            _gen_students_total(v)


def on_fetched_item(resource_name, response):
    if resource_name == 'classes':
        _gen_timestamp(response)
        _gen_students_total(response)


def on_inserted(resource_name, items):
    pass
    # if resource_name == 'attendances_tutors':
    #     blueprints._schedules.put_schedules()


def on_insert(resource_name, items):
    if resource_name == 'attendances_tutors':
        for i, v in enumerate(items):
            class_ = v.pop('class', None)

            lookup = {'_id': class_}
            class_, *_ = getitem_internal('classes', **lookup)

            try:
                lookup = {
                    'class': class_['_id'],
                    'module': class_['module'],
                    '_created': {'$gte': class_['startAtTs']},
                }
                attendance, *_ = getitem('attendances', **lookup)
            except NotFound:
                payload = {
                    'class': class_['_id'],
                    'module': class_['module']
                }
                attendance, *_ = post_internal('attendances', payload)

            try:
                lookup = {
                    'attendance': attendance['_id'],
                    'tutor': app.auth.get_request_auth_value(),
                    '_created': {'$gte': attendance['_created']},
                }
                *_, status = getitem('attendances_tutors', **lookup)
                pprint(status)
                if status == 200:
                    abort(422, description='exist')
            except NotFound:
                items[i].update({
                    'attendance': attendance['_id'],
                    'tutor': app.auth.get_request_auth_value(),
                })
