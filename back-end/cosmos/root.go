package cosmos

import (
	"fmt"
	"strconv"

	"github.com/notional-labs/multisignature-service/db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type TxHandler struct {
	Ctx client.Context
	Tx  db.Tx
}

func GetTxHandler(tx_id string) *TxHandler {
	txHandler := TxHandler{}

	// get tx
	tx := db.Tx{}
	tx.FindOne(tx_id)

	txHandler.Tx = tx

	// construct client context
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	ctx := client.Context{
		ChainID:           tx.Chain_id,
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
	}

	txHandler.Ctx = ctx

	return &txHandler
}

func (txHandler *TxHandler) ConstructMultisigTx() error {
	// parse TX
	parsedTx, err := txHandler.Ctx.TxConfig.TxJSONDecoder()([]byte(txHandler.Tx.Tx_body))
	if err != nil {
		return err
	}

	// construct tx builder
	txCfg := txHandler.Ctx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return err
	}

	// get signature list
	signs := db.FindAllSign(txHandler.Tx.Tx_id)

	// construct multisig
	/*
		for _, sign := range signs {
			// construct PubKey list
		}
	*/

	// signing
	sequenceNumber, err := strconv.ParseUint(txHandler.Tx.SequenceNumber, 10, 64)
	if err != nil {
		return err
	}

	for _, sign := range signs {
		parsedSigns, err := txHandler.Ctx.TxConfig.UnmarshalSignatureJSON([]byte(sign.Tx_id))
		if err != nil {
			return err
		}

		if txHandler.Ctx.ChainID == "" {
			return fmt.Errorf("ChainID is not set")
		}

		for _, sig := range parsedSigns {
			signingData := signing.SignerData{
				ChainID:  txHandler.Ctx.ChainID,
				Sequence: sequenceNumber,
			}

			err = signing.VerifySignature(sig.PubKey, signingData, sig.Data, txCfg.SignModeHandler(), txBuilder.GetTx())
			if err != nil {
				addr, _ := sdk.AccAddressFromHex(sig.PubKey.Address().String())
				return fmt.Errorf("couldn't verify signature for address %s", addr)
			}

			//if err := multisig.AddSignatureV2(multisigSig, sig, multisigPub.GetPubKeys()); err != nil {
			//	return err
			//}
		}

	}

	return nil
}

func (txHandler *TxHandler) BroadCastTx() error {
	return nil
}
