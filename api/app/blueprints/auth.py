from pprint import pprint
from datetime import timedelta
import jwt
import bcrypt
from flask_cors import CORS
from flask import current_app as app, Blueprint, request, abort
from eve.auth import requires_auth
from eve.utils import config
from eve.methods.get import getitem_internal
from eve_swagger import add_documentation
from shared.datetime import after_request_cache
from werkzeug.exceptions import NotFound
from eve.render import send_response
from eve.io.mongo.mongo import MongoJSONEncoder

blueprint = Blueprint('auth', __name__)
CORS(blueprint, max_age=timedelta(seconds=config.CACHE_EXPIRES))
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
        abort(422, description='username required')

    expected_username = data.get('username')
    if not expected_username:
        abort(422, description='username required')

    expected_username = expected_username.lower()
    expected_role = data.get('role')

    if not expected_role:
        abort(422, description='role required')

    lookup = {
        'username': expected_username,
        'role': expected_role[0]
    }
    try:
        user, *_ = getitem_internal(resource, **lookup)
    except NotFound:
        abort(404, description='user not found')

    if user.get('password'):
        password = data.get('password')
        try:
            if (not bcrypt.checkpw(password.encode(), user['password'].encode())):
                abort(404, description='password required or wrong password')
        except Exception as e:
            abort(404, description=str(e))

    payl = {config.ID_FIELD: user[config.ID_FIELD]}

    return send_response(None, ({
        'token': jwt.encode(payl, config.JWT_SECRET, json_encoder=MongoJSONEncoder)
    },))


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
        r = getitem_internal(resource, **lookup)

        return send_response('users', r)
    except Exception as e:
        abort(400, description=str(e))


@blueprint.after_request
def add_header(response):
    return after_request_cache(response)
