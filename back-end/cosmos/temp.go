package cosmos

import (
	"log"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	auth_tx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	bank_types "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func Temp() error {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := auth_tx.NewTxConfig(marshaler, auth_tx.DefaultSignModes)
	from_address, _ := sdk.AccAddressFromBech32("cosmos1p37nums898azmm9gc8yyxf8km97pu68p7pwgqv")
	ctx := client.Context{
		FromAddress:       from_address,
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
	}

	toAddr, err := sdk.AccAddressFromBech32("cosmos1p37nums898azmm9gc8yyxf8km97pu68p7pwgqv")
	if err != nil {
		return err
	}

	coins, err := sdk.ParseCoinsNormalized("2atom")
	if err != nil {
		return err
	}

	msg := bank_types.NewMsgSend(ctx.GetFromAddress(), toAddr, coins)

	log.Println(msg.String())

	return nil
}

func Temp1() error {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := auth_tx.NewTxConfig(marshaler, auth_tx.DefaultSignModes)
	amino := codec.NewLegacyAmino()
	from_address, _ := sdk.AccAddressFromBech32("cosmos1p37nums898azmm9gc8yyxf8km97pu68p7pwgqv")
	ctx := client.Context{
		FromAddress:       from_address,
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
		LegacyAmino:       amino,
	}

	coins, err := sdk.ParseCoinsNormalized("2atom")
	if err != nil {
		return err
	}

	br := rest.NewBaseReq(
		"cosmos1p37nums898azmm9gc8yyxf8km97pu68p7pwgqv",
		"",
		"cosmos-4",
		"5",
		"5",
		uint64(34),
		uint64(23),
		coins,
		nil,
		false,
	)

	toAddr, err := sdk.AccAddressFromBech32("cosmos1p37nums898azmm9gc8yyxf8km97pu68p7pwgqv")
	if err != nil {
		return err
	}

	msg := bank_types.NewMsgSend(ctx.GetFromAddress(), toAddr, coins)

	gasAdj, ok := ParseFloat64O(br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return nil
	}

	gasSetting, _ := flags.ParseGasSetting(br.Gas)

	txf := tx.Factory{}.
		WithFees("2atom").
		WithAccountNumber(br.AccountNumber).
		WithSequence(br.Sequence).
		WithGas(gasSetting.Gas).
		WithGasAdjustment(gasAdj).
		WithMemo(br.Memo).
		WithChainID(br.ChainID).
		WithSimulateAndExecute(br.Simulate).
		WithTxConfig(ctx.TxConfig).
		WithTimeoutHeight(br.TimeoutHeight)

	tx.GenerateTx(ctx, txf, msg)

	return nil
}

func ParseFloat64O(s string, defaultIfEmpty float64) (n float64, ok bool) {
	if len(s) == 0 {
		return defaultIfEmpty, true
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return n, false
	}

	return n, true
}
