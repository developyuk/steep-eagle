from pprint import pprint
from datetime import timedelta
import uuid

from flask_cors import CORS
from flask_mail import Mail, Message
from flask import current_app as app, Blueprint, request, abort
from eve.methods.patch import patch_internal
from eve.methods.get import getitem_internal
from eve.utils import config
from eve_swagger import add_documentation
from werkzeug.exceptions import NotFound
import bcrypt
from eve.render import send_response


blueprint = Blueprint('forgot_password', __name__)
CORS(blueprint, max_age=timedelta(seconds=config.CACHE_EXPIRES))
mail = Mail()

resource = 'users'
template = """Hi %s

ini password baru kamu ya %s"""


add_documentation({
    'paths': {'/forgot_password': {'post': {
        'summary': 'Reset password and send to email',
        'tags': ['User'],
    }}}
})


@blueprint.route('/forgot_password', methods=['POST'])
def forgot_password():
    mail.init_app(app)

    data = request.values or request.get_json()
    if not data:
        abort(422, description='username or email required')

    expected_username = data.get('username')
    if not expected_username:
        abort(422, description='username or email required')

    expected_role = data.get('role')
    if not expected_role:
        abort(422, description='role required')

    r = {
        'username': expected_username,
        'role': expected_role[0]
    }
    try:
        user, *_ = getitem_internal(resource, **r)
    except NotFound:
        r = {
            'email': expected_username,
            'role': expected_role[0]
        }
        try:
            user, *_ = getitem_internal(resource, **r)
        except NotFound:
            abort(404, description='username or email not found')

    if not user.get('email'):
        abort(422, description='email not found')

    new_password = str(uuid.uuid4())[:8]

    new_password_hashed = bcrypt.hashpw(
        new_password.encode(), bcrypt.gensalt())
    new_password_hashed = new_password_hashed.decode()

    lookup = {config.ID_FIELD: user[config.ID_FIELD]}
    payload = {'password': new_password_hashed}
    _ = patch_internal(resource, payload, **lookup)

    body = template % (expected_username, new_password)

    msg = Message("Reset Password", body=body, recipients=[user['email']])

    mail.send(msg)
    return send_response(None, ({'message': 'success'},))
