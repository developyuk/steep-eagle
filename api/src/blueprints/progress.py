from datetime import timedelta
from copy import deepcopy

from flask_cors import CORS
from flask import current_app as app, jsonify, Blueprint

from eve.methods import get, getitem
from eve.methods.post import post_internal
from eve.methods.put import put_internal
from eve.auth import requires_auth
from eve.render import render_json
from werkzeug.exceptions import NotFound

blueprint = Blueprint('progress', __name__)
CORS(blueprint, max_age=timedelta(days=10))


def _class_group_avg_rating(v):
    v['avg_rating'] = 0
    if len(v['_items']):
        avg = []
        for v2 in v['_items']:
            avg.append(v2['avg_rating'])

        v['avg_rating'] = sum(avg) / len(v['_items'])
    return v


def _class_avg_rating(v):
    attendances_students, *_ = get('attendances_students', **{
        'attendance_id': v['attendance']['id']
    })
    attendances_students = attendances_students['_items']
    avg = []

    v['avg_rating'] = 0
    if len(attendances_students) > 0:
        for v2 in attendances_students:
            avg2 = v2['rating_interaction'] + \
                v2['rating_cognition']+v2['rating_creativity']
            avg.append(avg2/3)

            v['avg_rating'] = sum(avg) / len(attendances_students)
    return v


def _group_total_class(v):
    v['total'] = len(v['_items'])
    return v


def _group_by_class(attendances):
    groups = []
    for v in attendances:
        class_id = v['attendance']['class_']['id']
        ii = [i for i, j in enumerate(groups) if j['class_id'] == class_id]
        item = {
            'id': v['id'],
            'attendance_id': v['attendance']['id'],
            'avg_rating': v['avg_rating'],
        }
        if not ii:
            group = {
                'class_id': class_id,
                'start_at': v['attendance']['class_']['start_at'],
                'finish_at': v['attendance']['class_']['finish_at'],
                'module': v['attendance']['module'],
                'branch': v['attendance']['class_']['branch'],
                '_items': [item]
            }
            groups.append(group)
        else:
            groups[ii[0]]['_items'].append(item)
    groups = map(_group_total_class, groups)
    groups = map(_class_group_avg_rating, groups)
    groups = list(groups)
    return groups


@blueprint.route('/progress/classes', methods=['GET'])
@requires_auth('/progress/classes')
def classes():
    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 9999
    app.config['DOMAIN']['attendances_tutors'].update({'embedded_fields': [
        'attendance',
        'attendance.class_',
        'attendance.class_.branch',
        'attendance.module',
    ]})

    lookup = {
        'tutor': app.auth.get_request_auth_value()
    }
    attendances, *_ = get('attendances_tutors', **lookup)
    attendances = attendances['_items']
    count_attendances = len(attendances)
    attendances = map(_class_avg_rating, attendances)

    attendances = _group_by_class(attendances)
    app.config = app_config_ori

    payload = {
        '_items': attendances,
        'meta': {
            'total': len(attendances),
            'total_item': count_attendances,
        }
    }
    try:
        lookup = {'key': 'progress_classes'}
        res, *_ = getitem('caches', **lookup)
        payload = {
            'value': render_json(payload)
        }
        res, *_ = put_internal('caches', payload, **{'id': res['id']})
    except NotFound:
        payload = {
            'key': 'progress_classes',
            'value': render_json(payload)
        }
        res, *_ = post_internal('caches', payload)

    # print(res)
    return jsonify(res)


def _student_group_avg_rating(v):
    v['avg_rating'] = 0
    if len(v['_items']):
        avg = []
        for v2 in v['_items']:
            avg.append(v2['avg_rating'])

        v['avg_rating'] = sum(avg) / len(v['_items'])
    return v


def _group_by_student(attendances):
    groups = []
    for v in attendances:
        student_id = v['student']['id']
        ii = [i for i, j in enumerate(groups) if j['student_id'] == student_id]
        item = {
            'id': v['id'],
            'attendance_id': v['attendance']['id'],
            'avg_rating': v['avg_rating'],
        }
        if not ii:
            group = {
                'student_id': student_id,
                'module': v['attendance']['module'],
                'branch': v['attendance']['class_']['branch'],
                '_items': [item]
            }
            groups.append(group)
        else:
            groups[ii[0]]['_items'].append(item)
    # groups = map(_group_total_class, groups)
    groups = map(_student_group_avg_rating, groups)
    groups = list(groups)

    return groups


def _student_avg_rating(v):
    v['avg_rating'] = (v['rating_interaction'] +
                       v['rating_cognition']+v['rating_creativity']) / 3

    return v


@blueprint.route('/progress/students', methods=['GET'])
@requires_auth('/progress/students')
def students():
    app_config_ori = deepcopy(app.config)
    app.config['PAGINATION_DEFAULT'] = 9999
    app.config['DOMAIN']['attendances_students'].update({'embedded_fields': [
        'attendance',
        'attendance.class_',
        'attendance.class_.branch',
        'attendance.module',
        'student',
    ]})
    lookup = {
        'tutor': app.auth.get_request_auth_value()
    }
    attendances, *_ = get('attendances_students', **lookup)
    attendances = attendances['_items']
    attendances = map(_student_avg_rating, attendances)
    attendances = list(attendances)
    count_attendances = len(attendances)

    attendances = _group_by_student(attendances)

    app.config = app_config_ori

    payload = {
        '_items': attendances,
        'meta': {
            'total': len(attendances),
            'total_item': count_attendances,
        }
    }
    try:
        lookup = {'key': 'progress_students'}
        res, *_ = getitem('caches', **lookup)
        payload = {
            'value': render_json(payload)
        }
        res, *_ = put_internal('caches', payload, **{'id': res['id']})
    except NotFound:
        payload = {
            'key': 'progress_students',
            'value': render_json(payload)
        }
        res, *_ = post_internal('caches', payload)

    # print(res)
    return jsonify(res)
