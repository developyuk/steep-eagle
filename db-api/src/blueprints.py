from flask import current_app as app, jsonify, Blueprint, request
from tables import Users, Exports
import pprint

blueprint = Blueprint('prefix_uri', __name__)


@blueprint.route('/sign', methods=['POST'])
def sign():
  data = request.values or request.get_json()

  try:
    username = data.get('username')
    user = app.data.driver.session.query(Users).filter(Users.username == username).first()
    return jsonify({'token': user.sign()})
  except AttributeError:
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': 'Please provide proper credentials',
        "code": 400
      }
    }), 400


@blueprint.route('/auth', methods=['GET'])
def auth():
  data = request.headers
  try:
    token = data.get('Authorization')
    token = token.split(' ')[1]
    user = Users.auth(token)
    return jsonify(user)
  except (AttributeError, TypeError):
    return jsonify({
      "_status": "ERR",
      "_error": {
        'message': 'Please provide proper credentials',
        "code": 400
      }
    }), 400
