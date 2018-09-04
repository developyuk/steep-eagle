from flask import current_app as app,Blueprint
from eve_swagger import add_documentation

blueprint = Blueprint('swagger', __name__)

add_documentation({
    'securityDefinitions': {
        "MyAuth": {"type": "apiKey", "name": "Authorization", "in": "header"}
    }
})

@blueprint.before_app_first_request
def init():
  for resource, rd in app.config['DOMAIN'].items():
      if (rd.get('disable_documentation')
              or resource.endswith('_versions')):
          continue

      methods = rd['resource_methods']
      url = '/%s' % rd['url']
      for method in methods:
          add_documentation({'paths': {url: {method.lower(): {
              "security": [{"MyAuth": []}],
              'produces': ['application/json']
          }}}})

      methods = rd['item_methods']
      item_id = '%sId' % rd['item_title'].lower()
      url = '/%s/{%s}' % (rd['url'], item_id)
      for method in methods:
          add_documentation({'paths': {url: {method.lower(): {
              "security": [{"MyAuth": []}],
              'produces': ['application/json']
          }}}})
