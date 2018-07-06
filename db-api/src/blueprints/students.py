from flask import current_app as app, jsonify, Blueprint
from eve.auth import requires_auth

blueprint = Blueprint('students', __name__)


@blueprint.route('/students', methods=['GET'])
@requires_auth('/students')
def students():
  return jsonify({})

@blueprint.after_request
def add_header(response):
  response.cache_control.max_age = app.config['CACHE_EXPIRES']
  response.cache_control.public = True
  response.cache_control.must_revalidate = True

  now = datetime.now()
  then = now + timedelta(days=app.config['CACHE_EXPIRES'])
  response.headers['Expires'] = then
  return response
