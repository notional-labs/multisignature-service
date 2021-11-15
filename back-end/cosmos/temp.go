package cosmos

import (
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func Temp() error {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	ctx := client.Context{
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
	}

	pkstr := `{"name":"multisig-0-1","type":"multi","address":"cosmos1z2mf7s005tqg3z7je7htkq5wguuruncjj89c5v","pubkey":"cosmospub1ytql0csgqgfzd666axrjzq6geytc3shllx83022aljarqe2zjnx405hszvs5e4quwtg2mwufmufzd666axrjzqle0w906f4ccprxnqt5mfpy89l62vzeqaw9lfs0s7z4qtcggpllpcpwylf6"}`

	var pk cryptotypes.PubKey
	err := ctx.Codec.UnmarshalInterfaceJSON([]byte(pkstr), &pk)
	log.Println(pk.String())

	return err
}
