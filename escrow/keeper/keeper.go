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

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey,
	supply expected.SupplyKeeper, joint expected.JoinKeeper) Keeper {
	return Keeper{
		cdc:          cdc,
		key:          key,
		supplyKeeper: supply,
		jointKeeper:  joint,
	}
}
