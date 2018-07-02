from eve.auth import TokenAuth
from tables import Users


class MyAuth(TokenAuth):
  def check_auth(self, token, allowed_roles, resource, method):
    # pprint.pprint(allowed_roles)
    # pprint.pprint(resource)
    # pprint.pprint(method)
    user = Users.auth(token)

    return user
