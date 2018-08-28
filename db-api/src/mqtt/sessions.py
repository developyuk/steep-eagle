import paho.mqtt.client as mqtt


def sessions_tutor_students(resource_name, items):

  if resource_name == 'sessions_students':
    # The callback for when the client receives a CONNACK response from the server.
    def on_connect(client, userdata, flags, rc):
      print("Connected with result code " + str(rc))

      # Subscribing in on_connect() means that if we lose the connection and
      # reconnect then subscriptions will be renewed.
      # client.subscribe("students")
      client.publish('students', 'tes')
      client.disconnect()

    client = mqtt.Client()
    client.on_connect = on_connect
    client.connect("127.0.0.1", 1883, 60)
    client.loop_forever()
