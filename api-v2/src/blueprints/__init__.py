from pytz import timezone, utc
from datetime import timedelta, datetime

from flask import current_app as app, abort
from eve.utils import parse_request
from eve.methods.common import ratelimit, epoch, pre_event, resolve_embedded_fields, \
    build_response_document, resource_link, document_link, last_updated

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

utc_now = utc.localize(datetime.utcnow())
wib_now = utc_now.astimezone(timezone("Asia/Jakarta"))


def onDay(day):
    now = utc_now
    return now + timedelta(days=(dow[day] - now.weekday() + 7) % 7)


from .auth import blueprint as _auth
from .schedules import blueprint as _schedules
from .students import blueprint as _students
from .forgot_password import blueprint as _forgot_password
from ._import import blueprint as _import
from .swagger import blueprint as _swagger
# from .calendar import blueprint as _calendar
# from .tutor_stats import blueprint as _tutor_stats
# from .progress import blueprint as _progress
