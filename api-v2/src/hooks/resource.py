from pprint import pprint

from pytz import timezone
from datetime import datetime
from shared.datetime import on_day


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
