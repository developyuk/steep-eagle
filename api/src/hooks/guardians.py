
import json

from eve.methods import post, getitem, delete
from eve.methods.post import post_internal
import sqlalchemy

_guardians = []


def update_student_guardians(guardians, student):
    for v2 in guardians:
        try:
            guardian, *_ = getitem('users', **{'username': v2['name']})
        except Exception:
            payload = {
                'username': v2['name'],
                'name': v2['name'],
                'email': v2['email'],
                'contact_no': v2['contact_no'],
                'role': 'guardian',
            }
            guardian, *_ = post_internal('users', payload)

        payload = {
            'student': student['id'],
            'guardian': guardian['id']
        }
        student_guardians, *_ = post_internal('student_guardians', payload)


def on_insert(resource_name, items):
    if resource_name == 'users':
        global _guardians
        for i, v in enumerate(items):
            if v['role'] == 'student':
                if v.get('guardians'):
                    _guardians[i] = json.loads(v.get('guardians'))
                    del v['guardians']


def on_inserted(resource_name, items):
    if resource_name == 'users':
        global _guardians
        for i, v in enumerate(items):
            if v['role'] == 'student':
                __guardians = _guardians[i]
                update_student_guardians(__guardians, v)


def on_update(resource_name, updates, original):
    if resource_name == 'users':
        if original['role'] == 'student':
            if updates.get('guardians'):
                _ = delete('student_guardians', **{'student': original['id']})
                _guardians = json.loads(updates.get('guardians'))
                update_student_guardians(_guardians, original)
                del updates['guardians']
