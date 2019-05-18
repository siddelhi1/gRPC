package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	// Populates database with dummy data
	var people []models.Person
	client, err := mongo.NewClient(CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(DBNAME)
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("person_data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValues, &people)
	// Insert people into DB
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}
	_, err = db.Collection(COLLECTION).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}
