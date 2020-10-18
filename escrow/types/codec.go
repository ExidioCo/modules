package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	ModuleCdc *codec.Codec
)

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreate{}, "exidioco/escrow/MsgCreate", nil)
	cdc.RegisterConcrete(MsgApprove{}, "exidioco/escrow/MsgApprove", nil)
}
