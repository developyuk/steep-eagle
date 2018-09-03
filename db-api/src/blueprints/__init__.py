from pytz import timezone, utc
from datetime import timedelta, datetime

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

utc_now = utc.localize(datetime.utcnow())
wib_now = utc_now.astimezone(timezone("Asia/Jakarta"))

from .auth import blueprint as bp_auth
from .swagger import blueprint as bp_swagger
from .schedules import blueprint as bp_schedules
from .students import blueprint as bp_students
from .calendar import blueprint as bp_calendar
from .tutor_stats import blueprint as bp_tutor_stats
