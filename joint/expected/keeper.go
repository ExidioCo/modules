package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
}
