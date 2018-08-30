from pprint import pprint

from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint, request, abort
from datetime import timedelta
from eve.auth import requires_auth
from eve.methods import getitem
from eve.utils import parse_request
import jwt

blueprint = Blueprint('auth', __name__)
CORS(blueprint, max_age=timedelta(days=10))
resource = 'users'


@blueprint.route('/sign', methods=['POST'])
def sign():
    data = request.values or request.get_json()

    try:
        username = data.get('username')
        username = username.lower()
        username = username[:-3]

        pprint(username)

        req = parse_request(resource)
        # req.show_deleted = True
        r = {'username': username}
        user = app.data.find_one(resource, req, **r)
        if (user['pass_']):
            password = data.get('password')
            if (password != user['pass_']):
                raise Exception('password required or wrong password')

        return jsonify({
            'token': jwt.encode({'id': user['id']}, app.config['JWT_SECRET'])
        })
    except Exception as e:
        abort(400, description=str(e))


@blueprint.route('/auth', methods=['GET'])
@requires_auth('/auth')
def auth():
    try:
        user, *_ = getitem(resource, {'id': app.auth.get_request_auth_value()})
        allowed_key = ('name', 'username', 'email', 'role', 'photo', 'id')
        user = filter(lambda v: v[0] in allowed_key, user.items())
        return jsonify(user)
    except Exception as e:
        abort(400, description=str(e))

# @blueprint.after_request
# def add_header(response):
#   response.cache_control.max_age = app.config['CACHE_EXPIRES']
#   response.cache_control.public = True
#   response.cache_control.must_revalidate = True
#
#   now = datetime.now()
#   then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
#   response.headers['Expires'] = then
#   return response
