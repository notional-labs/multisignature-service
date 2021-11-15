module github.com/notional-labs/multisignature-service

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.7.0 // indirect
	go.mongodb.org/mongo-driver v1.7.4
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
