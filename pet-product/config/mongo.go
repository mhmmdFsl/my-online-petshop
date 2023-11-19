package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func NewCollection(collenction string) *mongo.Collection {
	log.Printf("connecting to mongo\n")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed connecting to mongo: %v\n", err.Error())
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("failed ping mongodb: %v\n", err.Error())
	}
	c := client.Database("product").Collection(collenction)
	log.Printf("connected to mongodb\n")
	return c
}
