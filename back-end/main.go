package main

import (
	"github.com/notional-labs/multisignature-service/api_gateway"
	"github.com/notional-labs/multisignature-service/cosmos"
	"github.com/notional-labs/multisignature-service/db"
)

func main() {
	cosmos.Temp()
	db.InitDB()
	api_gateway.InitAPI()
}
