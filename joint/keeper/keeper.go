package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/expected"
)

type Keeper struct {
	cdc    *codec.Codec
	key    sdk.StoreKey
	supply expected.SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, supply expected.SupplyKeeper) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		supply: supply,
	}
}
