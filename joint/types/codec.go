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
	cdc.RegisterConcrete(MsgCreate{}, "exidioco/joint/MsgCreate", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "exidioco/joint/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgSend{}, "exidioco/joint/MsgSend", nil)
	cdc.RegisterConcrete(MsgApprove{}, "exidioco/joint/MsgApprove", nil)
}
