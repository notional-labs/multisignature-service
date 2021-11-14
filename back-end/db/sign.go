package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sign struct {
	Tx_id     string `bson:"tx_id" json:"tx_id"`
	Address   string `bson:"address" json:"address"`
	Sign_body string `bson:"sign_body" json:"sign_body"`
}

func (sign *Sign) UpdateOne() error {
	updateOpts := options.Update()
	updateOpts.SetUpsert(true)

	res, err := g_db.Collection("Sign").UpdateOne(
		g_ctx,
		bson.M{
			"tx_id":   sign.Tx_id,
			"address": sign.Address,
		},
		bson.M{
			"$set": sign,
		},
		updateOpts,
	)

	if err == nil {
		log.Printf("[SignCollection]: %d records added, %d records modified", res.UpsertedCount, res.ModifiedCount)
	}

	return err
}

func (sign *Sign) FindOne() error {
	findOpts := options.FindOne()

	res := g_db.Collection(("Sign")).FindOne(
		g_ctx,
		bson.M{
			"tx_id":   sign.Tx_id,
			"address": sign.Address,
		},
		findOpts,
	)

	if res.Err() == nil {
		res.Decode(sign)
	}

	return res.Err()
}
