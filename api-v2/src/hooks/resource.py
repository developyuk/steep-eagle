from pprint import pprint

from pytz import timezone, utc
from datetime import timedelta, datetime

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

utc_now = utc.localize(datetime.utcnow())
wib_now = utc_now.astimezone(timezone("Asia/Jakarta"))


def _on_day(day):
    now = utc_now
    return now + timedelta(days=(dow[day] - now.weekday() + 7) % 7)


def _gen_students_total(item):
    item['studentsTotal'] = 0


def _gen_timestamp(item):

    dt = datetime.strptime('%sT%s' % (
        _on_day(item['day']).date(), item['startAt']), '%Y-%m-%dT%H:%M')
    item['startAtTs'] = timezone('Asia/Jakarta').localize(dt)

    dt = datetime.strptime('%sT%s' % (
        _on_day(item['day']).date(), item['finishAt']), '%Y-%m-%dT%H:%M')
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
