from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS, cross_origin
from eve.auth import requires_auth
from datetime import timedelta, datetime

blueprint = Blueprint('tutor_stats', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/tutor_stats', methods=['GET'])
@requires_auth('/tutor_stats')
def tutor_stats():
  return jsonify({'_items': {
    'classes_sum': 0,
    'hours_sum': 0,
    'feedbacks_sum': 0,
    'ratings_avg': 0,
    'reviews_avg': 0,
    'attendances_avg': 0,
  }})

@blueprint.after_request
def add_header(response):
  response.cache_control.max_age = app.config['CACHE_EXPIRES']
  response.cache_control.public = True
  response.cache_control.must_revalidate = True

  now = datetime.now()
  then = now + timedelta(seconds=app.config['CACHE_EXPIRES'])
  response.headers['Expires'] = then
  return response
