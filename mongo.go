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

type Trick struct {
	ID     string `json:"id,omitempty" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Detail string `json:"detail" bson:"detail"`
	Level  string `json:"level" bson:"level"`
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

// trickInsert adds a record to our tricks collection in the hunter mongo-db.
// This is also used if the level of his trick has advanced
func (t Trick) trickInsert() {
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
	hunterCollection := hunterDatabase.Collection("tricks")
	hunterInsert, err := hunterCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: t.Name},
		{Key: "detail", Value: t.Detail},
		{Key: "level", Value: t.Level},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Insert: Full Trick: %v\n", hunterInsert)
}

// trainingSessionInsert only adds the name of the trick we practiced into our
// tricks collection in the hunter mongo-database
func (t Trick) trainingSessionInsert() {
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
	hunterCollection := hunterDatabase.Collection("tricks")
	hunterInsert, err := hunterCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: t.Name},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Insert: Training Session: %v\n", hunterInsert)
}

// readTricks takes a mongo database and collection string and returns
// a []Trick object of all of the documents within that collection
func readTricks(database, collection string) []Trick {
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

	hunterDatabase := client.Database(database)
	hunterCollection := hunterDatabase.Collection(collection)

	cursor, err := hunterCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// reading all of our documents into tricks at once
	var tricks []Trick
	if err = cursor.All(ctx, &tricks); err != nil {
		log.Fatal(err)
	}
	return tricks
}

// to be replaced with env var on deployment
func mongoPass() string {
	text, err := ioutil.ReadFile("mongo.txt")
	if err != nil {
		panic(err.Error())
	}
	return string(text)
}
