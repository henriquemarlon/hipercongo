package repository

import (
	"context"
	"github.com/henriquemarlon/hipercongo/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SensorLogRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorLogRepositoryMongo(client *mongo.Client, dbName string, collection string) *SensorRepositoryMongo {
	sensorsColl := client.Database(dbName).Collection(collection)
	return &SensorRepositoryMongo{
		Collection: sensorsColl,
	}
}

func (s *SensorRepositoryMongo) CreateSensorLog(sensorLog *entity.Log) error {
	result, err := s.Collection.InsertOne(context.TODO(), sensorLog)
	log.Printf("Inserting log into the MongoDB collection with id: %s", result)
	return err
}
