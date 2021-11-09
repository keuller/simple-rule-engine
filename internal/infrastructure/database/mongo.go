package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func NewConnection() (err error) {
	log.Println("[INFO] Open database connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27025/"))

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("[WARN] fail to ping the database server - %s \n", err.Error())
	} else {
		log.Printf("[INFO] Database connection established.")
	}

	db = client.Database("princing-service")
	log.Println("[INFO] Database connection has established...")
	return
}

func GetCollection(name string) *mongo.Collection {
	if db == nil {
		log.Println("[WARN] There is no connection active.")
		return nil
	}
	return db.Collection(name)
}
