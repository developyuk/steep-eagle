from datetime import timedelta, datetime, date
from pytz import timezone
from pprint import pprint
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from tables import Users, ClassesTs, Sessions
import humanize
from eve.auth import requires_auth
from eve.methods import get, getitem

blueprint = Blueprint('calendar', __name__)
CORS(blueprint, max_age=timedelta(days=10))

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))


@blueprint.route('/calendar', methods=['GET'])
@requires_auth('/calendar')
def schedules():
    date_from = date(2018, 1, 1)
    date_to = date.today() + timedelta(weeks=6*4)
    total_days = (date_to - date_from).days + 1  # inclusive 5 days
    date_range = map(lambda v: (
        date_from + timedelta(days=v)), range(total_days))

    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 999
    app.config['DOMAIN']['classes'].update({'embedded_fields': [
        'branch',
        'tutor',
        'module',
    ]})

    classes, *_ = get('classes')
    classes = classes['_items']
    # classes = app.data.driver.session.query(ClassesTs).all()

    user, *_ = getitem('users', {'id': app.auth.get_request_auth_value()})
    # user = app.data.driver.session.query(Users).get( app.auth.get_request_auth_value())

    def exclude_dummies_non_tester(v):
        if 'tester' not in user['username']:
            return 'dummies' not in v['module']['name']
        else:
            return True

    classes = filter(exclude_dummies_non_tester, classes)
    classes = list(classes)
    # classes.sort(key=lambda v: v['start_at_ts'])

    def date2str(date, time):
        return datetime.strptime('%sT%s' % (date, time), '%Y-%m-%dT%H:%M')

    calendars = []
    for v in date_range:
        for v2 in classes:
            if v.weekday() == dow[v2['day']]:
                calendars.append({
                    'id': v2['id'],
                    'title': '%s at %s by %s' % (v2['module']['name'].upper(), v2['branch']['name'].upper(), v2['tutor']['name'].upper()),
                    'allDay': False,
                    'start': date2str(v, v2['start_at']).strftime("%Y-%m-%d %H:%M:%S"),
                    'finish': date2str(v, v2['finish_at']).strftime("%Y-%m-%d %H:%M:%S")
                })

    # classes = []
    app.config = app_config_ori
    return jsonify({
        '_items': calendars,
        'meta': {'total': len(calendars)}
    })


@blueprint.after_request
def add_header(response):
    response.cache_control.max_age = app.config['CACHE_EXPIRES']
    response.cache_control.public = True
    response.cache_control.must_revalidate = True

    now = datetime.now()
    then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
    response.headers['Expires'] = then
    return response
