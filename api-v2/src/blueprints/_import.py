from datetime import timedelta
from pprint import pprint

from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint, request, abort
from eve.methods import post, delete, getitem
from eve.methods.post import post_internal
from eve.methods.put import put
from werkzeug.exceptions import NotFound
from eve_swagger import add_documentation

blueprint = Blueprint('import', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def allowed_key(allowed_key, item):
    return dict(item)


add_documentation({
    'paths': {'/import/branches': {'put': {
        'summary': 'Replaces a branches document',
        'tags': ['Import'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/import/branches', methods=['PUT'])
def _import_branches():
    data = request.json
    items = data['_items']

    def branch_convert(item):
        allowed_key = ('name', 'address')
        item = filter(lambda v: v[0] in allowed_key, item.items())
        item = dict(item)
        item = filter(lambda v: v[1] is not None, item.items())
        item = dict(item)
        return item
    items = map(branch_convert, items)
    items = list(items)

    status_all = True
    resource = 'branches'
    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items)
    pprint(response)
    status_all = status_all and response[3] == 201

    return jsonify({'status': status_all})


add_documentation({
    'paths': {'/import/modules': {'put': {
        'summary': 'Replaces a modules document',
        'tags': ['Import'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/import/modules', methods=['PUT'])
def _import_modules():
    data = request.json
    items = data['_items']

    def modules_convert(item):
        if item.get('image'):
            item['image'] = item['image'].split('/')[-1]

        allowed_key = ('name', 'image')
        item = filter(lambda v: v[0] in allowed_key, item.items())
        item = dict(item)
        item = filter(lambda v: v[1] is not None, item.items())
        item = dict(item)

        return item
    items = map(modules_convert, items)
    items = list(items)

    status_all = True
    resource = 'modules'

    app.config['DOMAIN'][resource]['schema']['image']['type'] = 'string'
    app.config['DOMAIN'][resource]['_media'] = []
    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items)
    pprint(response)
    status_all = status_all and response[3] == 201

    return jsonify({'status': status_all})


add_documentation({
    'paths': {'/import/users': {'put': {
        'summary': 'Replaces a users document',
        'tags': ['Import'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/import/users', methods=['PUT'])
def _import_users():
    data = request.json
    items = data['_items']

    def users_convert(item):

        # if item.get('photo'):
        #     item['photo'] = item['photo'].split('/')[-1]
        if item.get('pass_'):
            item['password'] = item['pass_']
        if item.get('is_deleted'):
            item['_deleted'] = item['is_deleted']
        if item.get('contact_no'):
            item['contact'] = item['contact_no']

        allowed_key = ('_deleted',
                       'password', 'username', 'email', 'role', 'password', 'name', 'dob', 'address', 'school', 'grade', 'contact')
        item = filter(lambda v: v[0] in allowed_key, item.items())
        item = dict(item)
        item = filter(lambda v: v[1] is not None, item.items())
        item = dict(item)

        return item
    items = map(users_convert, items)
    items = list(items)

    status_all = True
    resource = 'users'
    app.config['DOMAIN'][resource]['soft_delete'] = False
    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items, skip_validation=True)
    pprint(response)
    status_all = status_all and response[3] == 201

    return jsonify({'status': status_all})


add_documentation({
    'paths': {'/import/classes': {'put': {
        'summary': 'Replaces a classes document',
        'tags': ['Import'],
        "security": [{"JwtAuth": []}],
    }}}
})


@blueprint.route('/import/classes', methods=['PUT'])
def _import_classes():
    data = request.json
    items = data['_items']

    def classes_convert(item):
        if item.get('start_at'):
            item['startAt'] = item['start_at']
        if item.get('finish_at'):
            item['finishAt'] = item['finish_at']
        if item.get('module'):
            lookup = {
                'name': item['module']['name']
            }
            module, *_ = getitem('modules', **lookup)
            item['module'] = module['_id']
        if item.get('branch'):
            lookup = {
                'name': item['branch']['name']
            }
            branch, *_ = getitem('branches', **lookup)
            item['branch'] = branch['_id']
        if item.get('tutor'):
            lookup = {
                'name': item['tutor']['username']
            }
            try:
                tutor, *_ = getitem('users', **lookup)
                item['tutor'] = tutor['_id']
            except NotFound:
                item['tutor'] = 4529365169733632
                # item['tutor'] = None

        allowed_key = ('day', 'startAt', 'finishAt',
                       'module', 'branch', 'tutor')
        item = filter(lambda v: v[0] in allowed_key, item.items())
        item = dict(item)
        item = filter(lambda v: v[1] is not None, item.items())
        item = dict(item)

        return item
    items = map(classes_convert, items)
    items = list(items)
    pprint(items)

    status_all = True
    resource = 'classes'

    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items, skip_validation=True)
    pprint(response)
    status_all = status_all and response[3] == 201

    return jsonify({'status': status_all})
