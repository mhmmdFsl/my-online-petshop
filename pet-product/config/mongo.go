package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var Collecion *mongo.Collection

func init() {
	log.Printf("connecting to mongo\n")
	client, err :=  mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed connecting to mongo: %v\n", err.Error())
	}
	
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("failed ping mongodb: %v\n", err.Error())
	}
	Collecion = client.Database("product").Collection("pet_product")
	log.Printf("connected to mongodb\n")
}