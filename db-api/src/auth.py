from flask import current_app as app
from eve.auth import TokenAuth
from tables import Users
from pprint import pprint

class MyAuth(TokenAuth):
  def check_auth(self, token, allowed_roles, resource, method):
    # pprint(allowed_roles)
    # pprint(resource)
    # pprint(method)
    user = Users.auth(token)
    # pprint(user)
    self.set_request_auth_value(user['id'])
    # pprint(app.auth.get_request_auth_value())
    # current_app.auth.get_request_auth_value()

    return user
