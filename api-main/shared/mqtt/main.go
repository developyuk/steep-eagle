package mqtt

import (
  "github.com/davecgh/go-spew/spew"
  mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MqttSchedules(client mqtt.Client, msg mqtt.Message) {
  //var requestMsg Message
  //json.Unmarshal(msg.Payload(), &requestMsg)
  spew.Dump("TOPIC", msg.Topic())
  spew.Dump("MSG", string(msg.Payload()[:]))
}

func MqttStudents(client mqtt.Client, msg mqtt.Message) {
  //var requestMsg Message
  //json.Unmarshal(msg.Payload(), &requestMsg)
  spew.Dump("TOPIC", msg.Topic())
  spew.Dump("MSG", string(msg.Payload()[:]))
}
