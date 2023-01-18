package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI = "mongodb://localhost:27017"
)

func main() {
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal("cannot create mongo object")
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("cannot connect to database")
	}
	db := client.Database("coinmonsta")

	r, err := db.Collection("posts").InsertOne(ctx, bson.D{
		{Key: "jacek1", Value: "example text"}})
	if err != nil {
		log.Println("cannot save record", err)
		return
	}

	log.Println(r)
}
