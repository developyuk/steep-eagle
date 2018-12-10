from datetime import timedelta,datetime
from pprint import pprint

from flask import current_app as app, Blueprint, request
from eve.utils import config
from flask_cors import CORS
from eve.auth import requires_auth
from eve.render import send_response
from eve.methods.get import get_internal, getitem_internal
from shared.datetime import wib_tz, wib_now, utc_now, dow_list, after_request_cache
from werkzeug.exceptions import NotFound

blueprint = Blueprint('last_attendances', __name__)
CORS(blueprint, max_age=timedelta(seconds=config.CACHE_EXPIRES))


@blueprint.route('/classes/<string:cid>/last_attendances', methods=['GET'])
@requires_auth('/classes/cid/last_attendances')
def classes_last_attendances(cid):
    # utc_start_at = class_['start'].astimezone(timezone('UTC'))
    utc_start_at = utc_now.date()
    utc_start_at = datetime.combine(utc_start_at, datetime.min.time())
    try:
        lookup = {
            config.DATE_CREATED: {'$gte': utc_start_at},
            'class': cid
        }
        attendance, *_ = getitem_internal('attendances', **lookup)
        lookup = {'attendance': attendance[config.ID_FIELD]}
        r = get_internal('attendances_tutors', **lookup)
    except NotFound:
        r = ({config.ITEMS: []},)

    return send_response('attendances_tutors', r)
