package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var g_ctx context.Context
var g_db *mongo.Database

func InitDB() *mongo.Database {
	ctx := context.TODO()
	g_ctx = ctx
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	g_db = client.Database("MultiSignature")

	log.Println("Database is ready")

	return g_db
}
