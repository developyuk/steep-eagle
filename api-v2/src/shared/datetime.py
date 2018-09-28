
from pytz import timezone, utc
from datetime import timedelta, datetime

dow_list = 'monday tuesday wednesday thursday friday saturday sunday'.split()
dow = dict(zip(dow_list, range(7)))


utc_now = utc.localize(datetime.utcnow())
wib_tz = timezone("Asia/Jakarta")
wib_now = utc_now.astimezone(wib_tz)


def on_day(day):
    now = utc_now
    return now + timedelta(days=(dow[day] - now.weekday() + 7) % 7)
