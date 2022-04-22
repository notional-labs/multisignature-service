package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type KeyRingHandler struct {
	Ctx client.Context
}

// GetKeyringHandler keyringBackend is : test | os | local
func GetKeyringHandler(chainID string, keyringBackend string) (*KeyRingHandler, error) {
	KeyHandler := KeyRingHandler{}

	// construct client context
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	ctx := client.Context{
		ChainID:           chainID,
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
	}

	if keyringBackend != "" {
		kr, err := client.NewKeyringFromBackend(ctx, keyringBackend)
		if err != nil {
			return &KeyHandler, err
		}
		ctx = ctx.WithKeyring(kr)
	}
	KeyHandler.Ctx = ctx

	return &KeyHandler, nil
}

func (keyhandler *KeyRingHandler) addMultiSigKey(pubKeys []cryptotypes.PubKey, threshold int, txId string) error {
	ctx := keyhandler.Ctx
	if err := validateMultisigThreshold(threshold, len(pubKeys)); err != nil {
		return err
	}
	AminoPubsKey := multisig.NewLegacyAminoPubKey(threshold, pubKeys)
	_, err := ctx.Keyring.SaveMultisig(txId, AminoPubsKey)
	if err != nil {
		return err
	}
	return nil
}

func (keyhandler *KeyRingHandler) deleteMultiSigKey(txId string) error {
	if err := keyhandler.Ctx.Keyring.Delete(txId); err != nil {
		return err
	}
	return nil
}
