from datetime import timedelta
import jwt

from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint, request, abort
from eve.auth import requires_auth
from eve.utils import config
from eve.methods.get import getitem_internal
from eve_swagger import add_documentation

blueprint = Blueprint('auth', __name__)
CORS(blueprint, max_age=timedelta(days=10))
resource = 'users'


add_documentation({
    'paths': {'/sign': {'post': {
        'summary': 'Retrieves a User token',
        'tags': ['User'],
    }}}
})


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

    r = {
        'username': expected_username,
        'role': expected_role[0]
    }
    user, *_ = getitem_internal(resource, **r)

    if not user:
        abort(400, description='user not found')

    if user.get('pass_'):
        password = data.get('password')
        if (password != user['pass_']):
            abort(400, description='password required or wrong password')

    return jsonify({
        'token': jwt.encode({config.ID_FIELD: user[config.ID_FIELD]}, config.JWT_SECRET)
    })


add_documentation({
    'paths': {'/auth': {'get': {
        'summary': 'Retrieves a User document',
        'tags': ['User'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/auth', methods=['GET'])
@requires_auth('/auth')
def auth():
    try:
        lookup = {config.ID_FIELD: app.auth.get_request_auth_value()}
        user, *_ = getitem_internal(resource, **lookup)
        allowed_key = ('name', 'username', 'email', 'role',
                       'photo', config.ID_FIELD, config.ETAG)
        user = filter(lambda v: v[0] in allowed_key, user.items())
        user = dict(user)
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
