package main

import (
	"github.com/notional-labs/multisignature-service/api_gateway"
	"github.com/notional-labs/multisignature-service/db"
)

func main() {
	db := db.InitDB()
	api_gateway.InitAPI(db)
}
