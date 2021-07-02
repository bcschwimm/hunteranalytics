package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Behavior struct {
	ID    string `json:"id,omitempty" bson:"id"`
	Date  string `json:"date" bson:"date"`
	Crate int    `json:"crate" bson:"crate"`
	Notes string `json:"notes" bson:"notes"`
}

// behavior insert adds a record to our metrics collection in mongoDb
func (b Behavior) insert() {
	pass := mongoPass()
	client, err := mongo.NewClient(options.Client().ApplyURI(pass))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	hunterDatabase := client.Database("hunter")
	hunterCollection := hunterDatabase.Collection("metrics")
	hunterInsert, err := hunterCollection.InsertOne(ctx, bson.D{
		{Key: "date", Value: b.Date},
		{Key: "crate", Value: b.Crate},
		{Key: "notes", Value: b.Notes},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Insert: Behavior: %v\n", hunterInsert)
}

func mongoPass() string {
	text, err := ioutil.ReadFile("mongo.txt")
	if err != nil {
		panic(err.Error())
	}
	return string(text)
}
