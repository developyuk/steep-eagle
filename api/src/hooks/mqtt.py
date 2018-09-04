import os
import json
from copy import deepcopy

from flask import current_app as app
from customs import JSONEncoder
# import paho.mqtt.client as mqtt
import paho.mqtt.publish as mqtt_publish

from eve.methods import getitem

is_prod = os.environ['DEBUG'] == 0

tls = None
if is_prod:
    tls = {
        'ca_certs': "/etc/letsencrypt/live/mtutor.codingcamp.id/fullchain.pem",
        'certfile': "/etc/letsencrypt/live/mtutor.codingcamp.id/fullchain.pem",
        'keyfile': "/etc/letsencrypt/live/mtutor.codingcamp.id/privkey.pem",
        # 'tls_version': "<tls_version>",
        # 'ciphers': "<ciphers>"
    }
port = 1883
if is_prod:
    port = 8883


def publish(topic, payload):
    mqtt_publish.single(
        topic, payload, hostname=os.environ['HOST_MQTT'], tls=tls, port=port)


def on_deleted_item(resource_name, item):
    app_config_ori = deepcopy(app.config)

    if resource_name == 'attendances_students':

        r = {'id': item['student']}
        student, *_ = getitem('users', r)

        r = {'id': app.auth.get_request_auth_value()}
        user, *_ = getitem('users', r)

        m = json.dumps({
            'by': user,
            'on': "undoRateReview"
        }, cls=JSONEncoder)

        publish("students", m)

    if resource_name == 'attendances_tutors':
        app.config['DOMAIN']['classes'].update({'embedded_fields': [
            'module',
        ]})

        r = {'id': item['attendance']}
        attendance, *_ = getitem('attendances', r)

        r = {'id': attendance['class_']}
        klass, *_ = getitem('classes', r)

        r = {'id': app.auth.get_request_auth_value()}
        user, *_ = getitem('users', r)

        m = json.dumps({
            'on': "undo",
            'by': user,
            'class': klass
        }, cls=JSONEncoder)
        publish("schedules", m)

    app.config = app_config_ori


def on_inserted(resource_name, items):
    if resource_name == 'attendances_students':
        for i, item in enumerate(items):
            r = {'id': item['student']}
            student, *_ = getitem('users', r)

            r = {'id': app.auth.get_request_auth_value()}
            user, *_ = getitem('users', r)

            m = json.dumps({
                'by': user,
                'item': item,
                'on': "successRateReview"
            }, cls=JSONEncoder)
            publish("students", m)

    if resource_name == 'attendances_tutors':
        for i, item in enumerate(items):
            r = {'id': item['attendance']}
            attendance, *_ = getitem('attendances', r)

            r = {'id': app.auth.get_request_auth_value()}
            user, *_ = getitem('users', r)

            r = {'id': attendance['class_']}
            klass, *_ = getitem('classes', r)

            m = json.dumps({
                'on': "startYes",
                'by': user,
                'class': klass,
                'item': item,
            }, cls=JSONEncoder)
            publish("schedules", m)
