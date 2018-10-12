import random
from pprint import pprint

from eve.utils import config

# https://www.shareicon.net/pack/doraemon


def _gen_image(response, attr='image'):
    if isinstance(response, dict) and not response.get(attr):
        response[attr] = 'https://via.placeholder.com/48?text=image'

        if response.get('role') == 'tutor':
            response[attr] = 'https://images.weserv.nl/?il&w=64&url=dl.dropboxusercontent.com/s/xdy592qcgjvbbnp/Teacher.png'

        if response.get('role') == 'operation':
            response[attr] = 'https://images.weserv.nl/?il&w=64&url=www.shareicon.net/download/2016/01/16/249034_joy_256x256.png'

        if response.get('role') == 'student':
            my_list = ['www.shareicon.net/download/2016/01/16/249030_nobita_256x256.png',
                       'www.shareicon.net/download/2016/01/16/249031_takeshi_256x256.png',
                       'www.shareicon.net/download/2016/01/16/249036_shizuka_256x256.png',
                       'www.shareicon.net/download/2016/01/16/249037_suneo_256x256.png']
            response[attr] = 'https://images.weserv.nl/?il&w=64&url=%s' % random.choice(
                my_list)


def on_fetched_resource(resource_name, response):
    if resource_name == 'modules':
        for v in response[config.ITEMS]:
            _gen_image(v)

    if resource_name in ('users', 'students', 'tutors'):
        for v in response[config.ITEMS]:
            _gen_image(v, 'photo')

    if resource_name == 'classes_students':
        for v in response[config.ITEMS]:
            if v.get('student'):
                _gen_image(v['student'], 'photo')

    if resource_name == 'attendances_tutors':
        for v in response[config.ITEMS]:
            # print(v['attendance']['class_'])
            if type(v['attendance']) is dict and type(v['attendance']['class_']) is dict and type(v['attendance']['class_']['students'][0]) is dict:
                # if type(v['attendance']) is not int and type(v['attendance']['class_']['students'][0]) is dict:
                for v2 in v['attendance']['class_']['students']:
                    _gen_image(v2['student'], 'photo')


def on_fetched_item(resource_name, response):
    if resource_name == 'modules':
        _gen_image(response)

    if resource_name in ('users', 'students', 'tutors'):
        _gen_image(response, 'photo')
