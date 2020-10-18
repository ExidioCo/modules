package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	joint "github.com/exidioco/modules/joint/types"
)

type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
}

type JoinKeeper interface {
	GetAccount(ctx sdk.Context, identity uint64) (joint.Account, bool)
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to uint64, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, from uint64, to string, coins sdk.Coins) error
}
