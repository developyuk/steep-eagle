from datetime import timedelta, datetime, date
from copy import deepcopy

from flask import current_app as app, jsonify, Blueprint
from flask_cors import CORS
from eve.auth import requires_auth
from eve.methods import get

blueprint = Blueprint('tutor_stats', __name__)
CORS(blueprint, max_age=timedelta(days=10))

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))


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

    app.config['DOMAIN']['sessions_tutors'].update({'embedded_fields': [
        'session',
        'session.class_',
        # 'students.student'
    ]})
    classes, *_ = get('classes',
                      {'tutor_id': app.auth.get_request_auth_value()})
    classes = classes['_items']

    sessions_tutors, *_ = get('sessions_tutors',
                              {'tutor_id': app.auth.get_request_auth_value()})
    sessions_tutors = sessions_tutors['_items']

    sessions_students_feedback, *_ = get('sessions_students', {
        'tutor_id': app.auth.get_request_auth_value(),
        'feedback': "!=\"\""
    })
    sessions_students_feedback = sessions_students_feedback['_items']

    sessions_students_rating, *_ = get('sessions_students', {
        'tutor_id': app.auth.get_request_auth_value(),
        "rating_interaction": "!=0",
        "rating_cognition": "!=0",
        "rating_creativity": "!=0"
    })
    sessions_students_rating = sessions_students_rating['_items']

    classes_sum = 0
    sessions_students_sum = 0
    for v in date_range:
        for v2 in classes:
            if v.weekday() == dow[v2['day']]:
                classes_sum = classes_sum+1

                sessions_students_sum = sessions_students_sum + \
                    len(v2['students'])

    hours_sum = 0
    for v in sessions_tutors:
        finish_time = datetime.strptime(
            v['session']['class_']['finish_at'], '%H:%M')
        start_time = datetime.strptime(
            v['session']['class_']['start_at'], '%H:%M')
        hours_interval = (finish_time - start_time).total_seconds() / 3600
        hours_sum = hours_sum + hours_interval

    ratings_avg = 0
    if sessions_students_sum:
      ratings_avg = (len(sessions_students_rating) / sessions_students_sum) * 100

    reviews_avg = 0
    if sessions_students_sum:
      reviews_avg = (len(sessions_students_feedback) / sessions_students_sum) * 100

    attendances_avg = 0
    if classes_sum:
      attendances_avg = (len(sessions_tutors) / classes_sum) * 100

    app.config = app_config_ori
    return jsonify({'_items': {
        'classes_sum': len(sessions_tutors),
        'hours_sum': round(hours_sum, 0),
        'feedbacks_sum': len(sessions_students_feedback),
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
