
from datetime import timedelta
import uuid
import json

from flask import current_app as app, jsonify, Blueprint, request, abort
from flask_cors import CORS
from flask_mail import Mail, Message

from eve.methods.patch import patch_internal
from eve.utils import parse_request

blueprint = Blueprint('forgot_password', __name__)
CORS(blueprint, max_age=timedelta(days=10))
mail = Mail()

resource = 'users'
template = """Hi %s

ini password baru kamu ya %s"""


@blueprint.route('/forgot_password', methods=['POST'])
def forgot_password():
    mail.init_app(app)

    data = request.values or request.get_json()
    if not data:
        abort(422, description='username required')

    expected_username = data.get('username')
    if not expected_username:
        abort(422, description='username required')

    expected_role = data.get('role')
    if not expected_role:
        abort(422, description='role required')
    expected_role = json.loads(expected_role)

    req = parse_request(resource)
    # req.show_deleted = True
    r = {
        'username': expected_username,
        'role': expected_role[0]
    }
    user = app.data.find_one(resource, req, **r)

    if not user:
        abort(404, description='user not found')

    new_password = str(uuid.uuid4())[:8]

    _ = patch_internal(resource,
                       {'pass_': new_password}, **{'id': user['id']})

    body = template % (expected_username, new_password)

    msg = Message("Reset Password", body=body, recipients=[user['email']])
    mail.send(msg)
    return jsonify({})
