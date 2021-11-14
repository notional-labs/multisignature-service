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

func (tx *Tx) UpdateOne() error {
	updateOpts := options.Update()
	updateOpts.SetUpsert(true)

	res, err := g_db.Collection("Tx").UpdateOne(
		g_ctx,
		bson.M{
			"tx_id": tx.Tx_id,
		},
		bson.M{
			"$set": tx,
		},
		updateOpts,
	)

	if err == nil {
		log.Printf("[TxCollection]:%d records added, %d records modified", res.UpsertedCount, res.ModifiedCount)
	}

	return err
}

func (tx *Tx) FindOne(tx_id string) error {
	findOpts := options.FindOne()

	res := g_db.Collection(("Tx")).FindOne(
		g_ctx,
		bson.M{
			"tx_id": tx_id,
		},
		findOpts,
	)

	if res.Err() == nil {
		res.Decode(tx)
	}

	return res.Err()
}
