from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from tables import Users
from pprint import pprint


class MyAuth(TokenAuth):

    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials',
              response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        # pprint(allowed_roles)
        # pprint(resource)
        # pprint(method)
        user = Users.auth(token)

        if len(user):
            self.set_request_auth_value(user['id'])
        # pprint(app.auth.get_request_auth_value())
        # current_app.auth.get_request_auth_value()

        return len(user)
