

from pprint import pprint

import jwt
from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from eve.utils import config


def auth(token):
    try:
        data = jwt.decode(token, config.JWT_SECRET)
    except Exception as e:
        abort(400, description=str(e))

    return data


class JwtAuth(TokenAuth):
    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials', response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        user = auth(token)
        is_exist = len(user)

        if is_exist:
            self.set_request_auth_value(user[config.ID_FIELD])
        # pprint(app.auth.get_request_auth_value())

        return is_exist
