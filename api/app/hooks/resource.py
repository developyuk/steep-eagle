from pprint import pprint

from flask import current_app as app, abort
from eve.methods.get import getitem_internal
from eve.methods.patch import patch_internal
from eve.methods.post import post_internal
from eve.utils import config
from shared.datetime import wib_now, wib_tz, dow_list
from werkzeug.exceptions import NotFound


def on_insert(resource_name, items):
    if resource_name == 'attendances_tutors':
        for i, v in enumerate(items):
            class_ = v.pop('class', None)

            lookup = {config.ID_FIELD: class_}
            class_, *_ = getitem_internal('classes', **lookup)
            date = wib_now.date()

            # pprint('class_')
            # pprint(class_)

            try:
                lookup = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module'],
                    config.DATE_CREATED: {'$gte': class_['start'].astimezone(wib_tz).replace(
                        day=date.day, month=date.month, year=date.year)},
                }
                attendance, *_ = getitem_internal('attendances', **lookup)
            except NotFound:
                payload = {
                    'class': class_[config.ID_FIELD],
                    'module': class_['module']
                }
                attendance, *_ = post_internal('attendances', payload)

            # pprint('attendance')
            # pprint(attendance)

            try:
                lookup = {
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                    config.DATE_CREATED: {'$gte': attendance[config.DATE_CREATED]},
                }
                *_, status = getitem_internal('attendances_tutors', **lookup)

                if status == 200:
                    abort(422, description='exist')
            except NotFound:
                # pprint(items)
                items[i].update({
                    'attendance': attendance[config.ID_FIELD],
                    'tutor': app.auth.get_request_auth_value(),
                })
        # pprint(items)


def on_inserted(resource_name, items):
    if resource_name == 'classes-students':
        # pprint(items)
        payload = {
            'studentsTotal': len(items)
        }
        lookup = {
            config.ID_FIELD: items[0]['class']
        }
        patch_internal('classes', payload, **lookup)
