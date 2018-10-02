from datetime import timedelta
from pprint import pprint
from copy import deepcopy

from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint, request, abort
from eve.methods import post, delete, getitem
from eve.methods.post import post_internal
from eve.methods.put import put
from eve.utils import config
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
    items = data[config.ITEMS]

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
    items = data[config.ITEMS]

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

    domain_ori = deepcopy(app.config['DOMAIN'])
    app.config['DOMAIN'][resource]['schema']['image']['type'] = 'string'
    app.config['DOMAIN'][resource]['_media'] = []
    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items)
    pprint(response)
    status_all = status_all and response[3] == 201

    app.config['DOMAIN'] = domain_ori
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
    items = data[config.ITEMS]

    def users_convert(item):

        # if item.get('photo'):
        #     item['photo'] = item['photo'].split('/')[-1]
        if item.get('pass_'):
            item['password'] = item['pass_']
        if item.get('is_deleted'):
            item['_deleted'] = item['is_deleted']
        else:
            item['_deleted'] = False
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
    domain_ori = deepcopy(app.config['DOMAIN'])
    app.config['DOMAIN'][resource]['soft_delete'] = False
    try:
        _ = delete(resource)
    except NotFound:
        pass

    response = post_internal(resource, items, skip_validation=True)
    pprint(response)
    status_all = status_all and response[3] == 201

    app.config['DOMAIN'] = domain_ori
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
    items = data[config.ITEMS]

    domain_ori = deepcopy(app.config['DOMAIN'])
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
            item['module'] = module[config.ID_FIELD]
        if item.get('branch'):
            lookup = {
                'name': item['branch']['name']
            }
            branch, *_ = getitem('branches', **lookup)
            item['branch'] = branch[config.ID_FIELD]
        if item.get('tutor'):
            resource = 'users'
            app.config['DOMAIN'][resource]['soft_delete'] = False
            lookup = {
                'username': item['tutor']['username']
            }
            try:
                tutor, *_ = getitem(resource, **lookup)
                item['tutor'] = tutor[config.ID_FIELD]
            except NotFound:
                item['tutor'] = None

        if item.get('students'):
            resource = 'users'
            for v in item['students']:
                lookup = {
                    'username': v['student']['username']
                }
                try:
                    student, *_ = getitem(resource, **lookup)
                    v['student'] = student[config.ID_FIELD]
                except NotFound:
                    v['student'] = None

        allowed_key = ('day', 'startAt', 'finishAt',
                       'module', 'branch', 'tutor', 'students')
        item = filter(lambda v: v[0] in allowed_key, item.items())
        item = dict(item)
        item = filter(lambda v: v[1] is not None, item.items())
        item = dict(item)

        return item
    items = map(classes_convert, items)
    items = list(items)

    status_all = True
    resource = 'classes'
    resource_child = 'classes_students'

    try:
        _ = delete(resource)
        _ = delete(resource_child)
    except NotFound:
        pass

    for v in items:
        students = v.pop('students', None)
        class_, *response = post_internal(
            resource, v, skip_validation=True)
        status_all = status_all and response[2] == 201

        payload = [{'class': class_[config.ID_FIELD], 'student': v2['student']}
                   for v2 in students]
        response = post_internal(resource_child, payload)
        pprint(response)
        status_all = status_all and response[3] == 201

    app.config['DOMAIN'] = domain_ori
    return jsonify({'status': status_all})