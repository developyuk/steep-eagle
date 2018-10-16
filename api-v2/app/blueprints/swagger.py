import re
from flask import current_app as app, Blueprint
from eve_swagger import add_documentation

blueprint = Blueprint('swagger', __name__)

add_documentation({
    'securityDefinitions': {
        "JwtAuth": {"type": "apiKey", "name": "Authorization", "in": "header"}
    }
})


def _clear_regex(_str):
    return '/'.join(map(lambda x: re.sub('regex(.*):', '', x), _str.split('/')))


@blueprint.before_app_first_request
def init():
    for resource, rd in app.config['DOMAIN'].items():
        if (rd.get('disable_documentation')
                or resource.endswith('_versions')):
            continue

        methods = rd['resource_methods']
    #   url = '/%s' % rd['url']
        url = '/%s' % _clear_regex(rd['url'])
        for method in methods:
            add_documentation({'paths': {url: {method.lower(): {
                "security": [{"JwtAuth": []}],
                'produces': ['application/json']
            }}}})

        methods = rd['item_methods']
        item_id = '%sId' % rd['item_title'].lower()
        url = '/%s/{%s}' % (rd['url'], item_id)
        url = '/%s/{%s}' % (_clear_regex(rd['url']), item_id)
        for method in methods:
            add_documentation({'paths': {url: {method.lower(): {
                "security": [{"JwtAuth": []}],
                'produces': ['application/json']
            }}}})
