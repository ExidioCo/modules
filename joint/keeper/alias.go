package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/types"
)

func (k Keeper) Send(ctx sdk.Context, to sdk.AccAddress, coins sdk.Coins) error {
	return k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
}
