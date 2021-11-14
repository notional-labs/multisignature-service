package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tx struct {
	Tx_id   string `bson:"tx_id" json:"tx_id"`
	Tx_body string `bson:"tx_body" json:"tx_body"`
}

func createTxCollection() {
	collectionOpts := options.CreateCollection()
	err := g_db.CreateCollection(g_ctx, "Tx", collectionOpts)
	if err != nil {
		log.Fatal(err)
	}
}

func (tx *Tx) UpdateOne() error {
	updateOpts := options.Update()
	updateOpts.SetUpsert(true)

	res, err := g_db.Collection("Tx").UpdateOne(
		g_ctx,
		bson.M{
			"Tx_id": tx.Tx_id,
		},
		tx,
		updateOpts,
	)

	if err == nil {
		log.Printf("%d records changed", res.ModifiedCount)
	}

	return err
}

func (tx *Tx) FindOne() error {
	findOpts := options.FindOne()

	res := g_db.Collection(("Tx")).FindOne(
		g_ctx,
		bson.M{
			"Tx_id": tx.Tx_id,
		},
		findOpts,
	)

	if res.Err() == nil {
		res.Decode(tx)
	}

	return res.Err()
}
