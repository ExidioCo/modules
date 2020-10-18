package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SupplyKeeper interface {
	SendCoinsFromModuleToModule(ctx sdk.Context, from, to string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
}
