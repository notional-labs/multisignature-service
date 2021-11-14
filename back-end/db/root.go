package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database("MultiSignature")
}
