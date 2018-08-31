import random

# https://www.shareicon.net/pack/doraemon
def image_empty(response, attr='image'):
    if not response[attr]:
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
        for v in response['_items']:
            image_empty(v)

    if resource_name == 'users':
        for v in response['_items']:
            image_empty(v, 'photo')

    if resource_name == 'sessions_tutors':
        for v in response['_items']:
            if type(v['session']['class_']['students'][0]) is dict:
                for v2 in v['session']['class_']['students']:
                    image_empty(v2['student'], 'photo')


def on_fetched_item(resource_name, response):
    if resource_name == 'modules':
        image_empty(response)

    if resource_name == 'users':
        image_empty(response, 'photo')
