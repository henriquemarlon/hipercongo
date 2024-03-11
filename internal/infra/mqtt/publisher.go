package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/henriquemarlon/hipercongo/internal/domain/entity"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type PublisherMQTTRepository struct {
	Topic  string
	QoS    byte
	Client MQTT.Client
}

func NewPublisherMQTTRepository(topic string, qos byte, client MQTT.Client) *PublisherMQTTRepository {
	return &PublisherMQTTRepository{
		Topic:  topic,
		QoS:    qos,
		Client: client,
	}
}

func (s *PublisherMQTTRepository) Publish(data *entity.Log) error {
	for {
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return err
		}
		token := s.Client.Publish(s.Topic, s.QoS, false, string(payload))
		log.Println("Published: ", string(payload))
		token.Wait()
		time.Sleep(3 * time.Second)
	}
}