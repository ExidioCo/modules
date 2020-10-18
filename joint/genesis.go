package joint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	var count uint64
	for _, account := range state {
		if account.Identity > count {
			count = account.Identity
		}

		k.SetAccount(ctx, account)
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetAccounts(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, account := range state {
		if err := account.Validate(); err != nil {
			return err
		}
	}

	accounts := make(map[uint64]bool)
	for _, account := range state {
		if accounts[account.Identity] {
			return fmt.Errorf("duplicate account identity %d", account.Identity)
		}

		accounts[account.Identity] = true
	}

	return nil
}
