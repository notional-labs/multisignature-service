package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var g_ctx context.Context
var g_db *mongo.Database

func InitDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	g_ctx = ctx
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")

	defer cancel()
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

	createTxCollection()

	return g_db
}

func createSignCollection() {

}
