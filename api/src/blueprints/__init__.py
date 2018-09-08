from pytz import timezone, utc
from datetime import timedelta, datetime

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

utc_now = utc.localize(datetime.utcnow())
wib_now = utc_now.astimezone(timezone("Asia/Jakarta"))

from .auth import blueprint as _auth
from .swagger import blueprint as _swagger
from .schedules import blueprint as _schedules
from .students import blueprint as _students
from .calendar import blueprint as _calendar
from .tutor_stats import blueprint as _tutor_stats
from .forgot_password import blueprint as _forgot_password
from .progress import blueprint as _progress
