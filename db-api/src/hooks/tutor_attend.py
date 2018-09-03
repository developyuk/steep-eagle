from eve.methods import getitem, post, deleteitem
from eve.methods.post import post_internal
from eve.methods.delete import deleteitem_internal
from blueprints import utc_now
from werkzeug.exceptions import NotFound


def on_deleted_item(resource_name, item):
    if resource_name == 'attendances_tutors':
        # r = {'id': item['attendance']}
        # attendance, *_ = deleteitem_internal( 'attendances', *r)
        try:
            other_attendances_tutors, *_ = getitem('attendances_tutors', {
                'attendance': item['attendance']
            })
        except NotFound as e:
            attendance, *_ = deleteitem_internal(
                'attendances', id=item['attendance'])


def on_insert(resource_name, items):
    if resource_name == 'attendances_tutors':
        for i, item in enumerate(items):

            _class, *_ = getitem('classes', {
                'id': item['class_']
            })
            del item['class_']

            attendance = None
            try:
                attendance, *_ = getitem('attendances', {
                    '_created': '>="%s"' % utc_now.date().strftime('%Y-%m-%d'),
                    'class_': _class['id']
                })
            except NotFound as e:
                payl = {
                    'class_': _class['id'],
                    'module': _class['module'],
                }
                attendance, *_ = post_internal('attendances', payl)

            items[i]['attendance'] = attendance['id']
