package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/escrow/expected"
)

type Keeper struct {
	cdc          *codec.Codec
	key          sdk.StoreKey
	supplyKeeper expected.SupplyKeeper
	jointKeeper  expected.JoinKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}
