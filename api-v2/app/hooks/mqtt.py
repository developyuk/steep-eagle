import os
import json
from pprint import pprint


from eve.io.mongo.mongo import MongoJSONEncoder
from flask import current_app as app
# import paho.mqtt.client as mqtt
import paho.mqtt.publish as mqtt_publish


port = 1883
project_name = 'steep-eagle'


def publish(topic, payload):
    #     pprint(os.environ['HOST_MQTT'], topic, payload, port)
    mqtt_publish.single(
        topic, payload, hostname=os.environ['HOST_MQTT'], port=port)


def on_deleted_item(resource_name, item):
    if resource_name == 'attendances_students':
        m = json.dumps({
            'by': app.auth.get_request_auth_value(),
            'update': True,
        }, cls=MongoJSONEncoder)
        publish(f"{project_name}.students", m)

    if resource_name == 'attendances_tutors':
        m = json.dumps({
            'by': app.auth.get_request_auth_value(),
            'update': True,
        }, cls=MongoJSONEncoder)
        publish(f"{project_name}.schedules", m)


def on_inserted(resource_name, items):
    if resource_name == 'attendances-students':
        m = json.dumps({
            'by': app.auth.get_request_auth_value(),
            'update': True,
        }, cls=MongoJSONEncoder)
        publish(f"{project_name}.students", m)

    if resource_name == 'attendances_tutors':
        m = json.dumps({
            'by': app.auth.get_request_auth_value(),
            'update': True,
        }, cls=MongoJSONEncoder)
        publish(f"{project_name}.schedules", m)
