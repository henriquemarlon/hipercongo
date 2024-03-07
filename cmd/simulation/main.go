package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/henriquemarlon/hipercongo/internal/domain/entity"
	"github.com/henriquemarlon/hipercongo/internal/infra/repository"
	"github.com/henriquemarlon/hipercongo/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	options := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=orbit-city", os.Getenv("MONGODB_ATLAS_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		log.Printf("Identicating sensor: %v", sensor)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s:%s", os.Getenv("BROKER_TLS_URL"), os.Getenv("BROKER_PORT"))).SetUsername(os.Getenv("BROKER_USERNAME")).SetPassword(os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			for {
				payload, err := entity.NewSensorPayload(
					sensor.ID,
					sensor.Params,
					time.Now(),
				)
				if err != nil {
					log.Fatalf("Failed to create payload: %v", err)
				}

				jsonPayload, err := json.Marshal(payload)
				if err != nil {
					log.Println("Error converting to JSON:", err)
				}
				token := client.Publish("sensors", 1, false, string(jsonPayload))
				log.Printf("Published: %s, on topic: %s", string(jsonPayload), "sensors")
				token.Wait()
				time.Sleep(360 * time.Second)
			}
		}(sensor)
	}
	wg.Wait()
}
