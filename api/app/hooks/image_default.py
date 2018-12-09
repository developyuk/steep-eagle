import random
from pprint import pprint

# https://www.shareicon.net/pack/doraemon

default_photo = {
    'tutor': 'teacher.png',
    'operation': 'joy.png',
    'student': ['nobita.png', 'takeshi.png', 'shizuka.png', 'suneo.png']
}


def gen_image(resource_name, item):
    if resource_name == 'modules':
        return 'ph.gif'

    if resource_name == 'users':
        role = item['role']
        if isinstance(default_photo[role], list):
            return random.choice(default_photo[role])
        else:
            return default_photo[role]


def on_insert(resource_name, items):
    if resource_name == 'modules':
        for v in items:
            v['image'] = gen_image(resource_name, v)

    if resource_name == 'users':
        for v in items:
            v['photo'] = gen_image(resource_name, v)
