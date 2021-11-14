package cosmos

import (
	"log"

	"github.com/notional-labs/multisignature-service/db"

	"github.com/cosmos/cosmos-sdk/client"
	cosmos_tx "github.com/cosmos/cosmos-sdk/client/tx"
)

func ConstructMultisigTx(tx_id string) error {
	// construct client context
	ctx := client.Context{}

	// parse TX
	tx := db.Tx{}
	tx.FindOne(tx_id)

	parsedTx, err := ctx.TxConfig.TxJSONDecoder()([]byte(tx.Tx_body))
	if err != nil {
		log.Fatal(err)
	}

	// construct tx factory
	txFactory := cosmos_tx.Factory{}

	// construct tx builder
	txCfg := ctx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return err
	}

	// get signature list
	signs := db.FindAllSign(tx_id)

}

func BroadCastTx() {

}
