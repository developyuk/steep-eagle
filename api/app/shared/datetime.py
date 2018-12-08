
from pytz import timezone, utc
from datetime import timedelta, datetime
from eve.utils import config

dow_list = 'monday tuesday wednesday thursday friday saturday sunday'.split()
dow = dict(zip(dow_list, range(7)))


utc_now = utc.localize(datetime.utcnow())
wib_tz = timezone("Asia/Jakarta")
wib_now = utc_now.astimezone(wib_tz)


def on_day(day):
    return utc_now + timedelta(days=(dow[day] - utc_now.weekday() + 7) % 7)


def after_request_cache(response):
    response.cache_control.max_age = config.CACHE_EXPIRES
    response.cache_control.public = True
    response.cache_control.must_revalidate = True

    then = utc_now + timedelta(seconds=config.CACHE_EXPIRES)
    response.headers['Expires'] = then.strftime("%a, %d %b %Y %H:%M:%S GMT")
    return response
