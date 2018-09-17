from datetime import timedelta, datetime, date
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get
from . import dow

blueprint = Blueprint('tutor_stats', __name__)
CORS(blueprint, max_age=timedelta(days=10))


@blueprint.route('/tutor_stats', methods=['GET'])
@requires_auth('/tutor_stats')
def tutor_stats():
    app_config_ori = deepcopy(app.config)

    date_from = date(2018, 3, 10)
    date_to = date.today()
    total_days = (date_to - date_from).days + 1  # inclusive 5 days
    date_range = map(lambda v: (
        date_from + timedelta(days=v)), range(total_days))

    app.config['DOMAIN']['classes'].update({'embedded_fields': [
        'students',
    ]})

    app.config['DOMAIN']['attendances_tutors'].update({'embedded_fields': [
        'attendance',
        'attendance.class_',
        # 'students.student'
    ]})
    classes, *_ = get('classes',
                      **{'tutor_id': app.auth.get_request_auth_value()})
    classes = classes['_items']

    attendances_tutors, *_ = get('attendances_tutors',
                                 **{'tutor_id': app.auth.get_request_auth_value()})
    attendances_tutors = attendances_tutors['_items']

    attendances_students_feedback, *_ = get('attendances_students', **{
        'tutor_id': app.auth.get_request_auth_value(),
        'feedback': "!=\"\""
    })
    attendances_students_feedback = attendances_students_feedback['_items']

    attendances_students_rating, *_ = get('attendances_students', **{
        'tutor_id': app.auth.get_request_auth_value(),
        "rating_interaction": "!=0",
        "rating_cognition": "!=0",
        "rating_creativity": "!=0"
    })
    attendances_students_rating = attendances_students_rating['_items']

    classes_sum = 0
    attendances_students_sum = 0
    for v in date_range:
        for v2 in classes:
            if v.weekday() == dow[v2['day']]:
                classes_sum = classes_sum+1

                attendances_students_sum = attendances_students_sum + \
                    len(v2['students'])

    hours_sum = 0
    for v in attendances_tutors:
        finish_time = datetime.strptime(
            v['attendance']['class_']['finish_at'], '%H:%M')
        start_time = datetime.strptime(
            v['attendance']['class_']['start_at'], '%H:%M')
        hours_interval = (finish_time - start_time).total_seconds() / 3600
        hours_sum = hours_sum + hours_interval

    ratings_avg = 0
    if attendances_students_sum:
        ratings_avg = (len(attendances_students_rating) /
                       attendances_students_sum) * 100

    reviews_avg = 0
    if attendances_students_sum:
        reviews_avg = (len(attendances_students_feedback) /
                       attendances_students_sum) * 100

    attendances_avg = 0
    if classes_sum:
        attendances_avg = (len(attendances_tutors) / classes_sum) * 100

    app.config = app_config_ori
    return jsonify({'_items': {
        'classes_sum': len(attendances_tutors),
        'hours_sum': round(hours_sum, 0),
        'feedbacks_sum': len(attendances_students_feedback),
        'ratings_avg': round(ratings_avg, 2),
        'reviews_avg': round(reviews_avg, 2),
        'attendances_avg': round(attendances_avg, 2),
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
