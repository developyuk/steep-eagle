from pprint import pprint
from datetime import timedelta
import uuid

from flask_cors import CORS
from flask_mail import Mail, Message
from flask import current_app as app, jsonify, Blueprint, request, abort
from eve.methods.patch import patch_internal
from eve.methods.get import getitem_internal
from eve.utils import config
from eve_swagger import add_documentation


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
    user, *_ = getitem_internal(resource, **r)

    if not user:
        r = {
            'email': expected_username,
            'role': expected_role[0]
        }
        user, *_ = getitem_internal(resource, **r)

    if not user:
        abort(404, description='username or email not found')
    else:
        if not user['email']:
            abort(422, description='username or email found, but no email')

    new_password = str(uuid.uuid4())[:8]

    _ = patch_internal(
        resource, {'password': new_password}, **{config.ID_FIELD: user[config.ID_FIELD]})

    body = template % (expected_username, new_password)

    msg = Message("Reset Password", body=body, recipients=[user['email']])

    mail.send(msg)
    return jsonify({})
