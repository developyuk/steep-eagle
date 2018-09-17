from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint, request, abort
from datetime import timedelta
from eve.auth import requires_auth
from eve.methods import getitem
from . import getitem_internal
import jwt
import json

blueprint = Blueprint('auth', __name__)
CORS(blueprint, max_age=timedelta(days=10))
resource = 'users'


@blueprint.route('/sign', methods=['POST'])
def sign():
    data = request.values or request.get_json()

    if not data:
        abort(400, description='username required')

    expected_username = data.get('username')
    if not expected_username:
        abort(400, description='username required')

    expected_username = expected_username.lower()
    expected_role = data.get('role')

    if not expected_role:
        abort(400, description='role required')

    expected_role = json.loads(expected_role)
    # req.show_deleted = True
    r = {
        'username': expected_username,
        'role': expected_role[0]
    }
    user, *_ = getitem_internal(resource, **r)

    if not user:
        abort(400, description='user not found')

    if user['pass_']:
        password = data.get('password')
        if (password != user['pass_']):
            abort(400, description='password required or wrong password')

    return jsonify({
        'token': jwt.encode({'id': user['id']}, app.config['JWT_SECRET'])
    })


@blueprint.route('/auth', methods=['GET'])
@requires_auth('/auth')
def auth():
    try:
        user, *_ = getitem(resource,
                           **{'id': app.auth.get_request_auth_value()})
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
