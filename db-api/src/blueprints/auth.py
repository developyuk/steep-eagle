from flask import current_app as app, jsonify, Blueprint, request
from flask_cors import CORS, cross_origin
from tables import Users
from pprint import pprint
from datetime import timedelta
from eve.auth import requires_auth

blueprint = Blueprint('auth', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/sign', methods=['POST'])
def sign():
  data = request.values or request.get_json()

  try:
    username = data.get('username')
    user = app.data.driver.session.query(Users).filter(Users.username == username).first()
    return jsonify({'token': user.sign()})
  except Exception as e:
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': str(e),
        "code": 400
      }
    }), 400


@blueprint.route('/auth', methods=['GET'])
@requires_auth('/auth')
def auth():
  data = request.headers
  try:
    token = data.get('Authorization')
    token = token.split(' ')[1]
    user = Users.auth(token)
    user = app.data.driver.session.query(Users).get(user['id'])
    return jsonify({
      'id': user.id,
      'username': user.username,
      'email': user.email,
      'role': user.role,
      'photo': user.profile.photo,
    })
  except Exception as e:
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': str(e),
        "code": 400
      }
    }), 400
