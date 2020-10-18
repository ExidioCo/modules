package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/escrow/types"
	joint "github.com/exidioco/modules/joint/types"
)

func (k Keeper) GetAccount(ctx sdk.Context, identity uint64) (joint.Account, bool) {
	return k.jointKeeper.GetAccount(ctx, identity)
}

func (k Keeper) Lock(ctx sdk.Context, from uint64, coins sdk.Coins) error {
	return k.jointKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins)
}

func (k Keeper) Unlock(ctx sdk.Context, to uint64, coins sdk.Coins) error {
	return k.jointKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
}

func (k Keeper) SendCoinsToAccount(ctx sdk.Context, to sdk.AccAddress, coins sdk.Coins) error {
	return k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
}
