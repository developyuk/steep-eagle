from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from eve.auth import requires_auth
from datetime import timedelta, datetime

blueprint = Blueprint('tutor_stats', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/tutor/id/stats', methods=['GET'])
@requires_auth('/tutor/id/stats')
def tutor_stats():
  return jsonify({})
