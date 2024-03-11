package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SubscriberMQTTRepository struct {
	Topic  string
	QoS    byte
	Client MQTT.Client
}

func NewSubscriberMQTTRepository(topic string, qos byte, client MQTT.Client) *SubscriberMQTTRepository {
	return &SubscriberMQTTRepository{
		Topic:  topic,
		QoS:    qos,
		Client: client,
	}
}

func (s *SubscriberMQTTRepository) Subscribe(callback MQTT.MessageHandler) error {
	for {
		token := s.Client.Subscribe(s.Topic, s.QoS, callback)
		token.Wait()
	}
}