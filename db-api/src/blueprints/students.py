from flask import current_app as app, jsonify, Blueprint
from eve.auth import requires_auth

blueprint = Blueprint('students', __name__)


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
  return jsonify({})
