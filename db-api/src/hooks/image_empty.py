from pprint import pprint


def image_empty(response, attr='image'):
    if not response[attr]:
        response[attr] = 'https://via.placeholder.com/48?text=image'


def on_fetched_resource(resource_name, response):
    if resource_name == 'modules':
        for v in response['_items']:
            image_empty(v)

    if resource_name == 'users':
        for v in response['_items']:
            image_empty(v, 'photo')


def on_fetched_item(resource_name, response):
    if resource_name == 'modules':
        image_empty(response)
    if resource_name == 'users':
        image_empty(response, 'photo')
