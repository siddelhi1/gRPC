package dao

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// Connect establish a connection to database

func init() {
	client, err := mongo.NewClient(CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}
