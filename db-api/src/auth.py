from eve.auth import TokenAuth
from tables import Users
from pprint import pprint

class MyAuth(TokenAuth):
  def check_auth(self, token, allowed_roles, resource, method):
    # pprint(allowed_roles)
    # pprint(resource)
    # pprint(method)
    user = Users.auth(token)

    return user
