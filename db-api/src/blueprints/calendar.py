from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from tables import Users, ClassesTs, Sessions
from datetime import timedelta, datetime
import humanize
from eve.auth import requires_auth
from pytz import timezone
from pprint import pprint

blueprint = Blueprint('calendar', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/calendar', methods=['GET'])
@requires_auth('/calendar')
def schedules():
    classes = app.data.driver.session.query(ClassesTs) \
        .all()

    user = app.data.driver.session.query(Users).get(
        app.auth.get_request_auth_value())

    def exclude_dummies_non_tester(v):
        if 'tester' not in user.username:
            return 'dummies' not in v.module.name
        else:
            return True

    classes = filter(lambda v: v.finish_at_ts > (
        datetime.now(timezone('UTC')) - timedelta(hours=2)), classes)
    classes = filter(lambda v: v.finish_at_ts.date() < (datetime.now(timezone('UTC')) + timedelta(days=5)).date(),
                     classes)
    classes = filter(exclude_dummies_non_tester, classes)
    classes.sort(key=lambda v: v.start_at_ts)

    def parse(v):
        w = {
            'id': v.id,
            'title': '%s at %s by %s' % (v.module.name, v.branch.name, v.tutor.name),
            'allDay': False,
            'start': v.start_at_ts.replace(tzinfo=None),
            'finish': v.finish_at_ts.replace(tzinfo=None)
        }

        return w

    classes = map(parse, classes)
    # classes = []
    return jsonify({
        '_items': classes,
        'meta': {'total': len(classes)}
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
