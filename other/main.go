package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var result struct {
	_id  string
	name string
}

func main() {

	//Query parameter for GET  result

	// Below functional shall be used to connect with Mongo DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	//mongodb://10.5.74.32:27017

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("Blockchain").Collection("profile")
	filter := bson.M{"_id": "1"}
	err = collection.FindOne(ctx, filter).Decode(&result)
	//	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

}
