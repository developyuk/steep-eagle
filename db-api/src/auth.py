from flask import Blueprint, request, current_app as app, jsonify
from eve.auth import TokenAuth
from tables import Users

blueprint = Blueprint('prefix_uri', __name__)


class MyAuth(TokenAuth):
  def check_auth(self, token, allowed_roles, resource, method):
    user = Users.auth(token)

    return user


@blueprint.route('/sign', methods=['POST'])
def sign():
  data = request.values or request.get_json()
  userid = data.get('userid')
  if not userid:
    return None
  else:
    user = app.data.driver.session.query(Users).filter(Users.username == userid).first()

    return jsonify({'token': user.sign()})


@blueprint.route('/auth', methods=['GET'])
def auth():
  data = request.headers
  token = data.get('Authorization')
  token = token.split(' ')[1]
  if not token:
    return None
  else:
    user = Users.auth(token)
    user = app.data.driver.session.query(Users).get(user['id'])

    return jsonify({'username': user.username})
