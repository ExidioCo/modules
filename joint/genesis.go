package joint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	var accounts uint64
	for _, account := range state.Accounts {
		if account.Identity > accounts {
			accounts = account.Identity
		}

		k.SetAccount(ctx, account)
	}

	k.SetAccountsCount(ctx, accounts)

	var transfers uint64
	for _, transfer := range state.Transfers {
		if transfer.Identity > transfers {
			transfers = transfer.Identity
		}

		k.SetTransfer(ctx, transfer)
	}

	k.SetTransfersCount(ctx, transfers)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(
		k.GetAccounts(ctx),
		k.GetTransfers(ctx),
	)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, account := range state.Accounts {
		if err := account.Validate(); err != nil {
			return err
		}
	}

	accounts := make(map[uint64]bool)
	for _, account := range state.Accounts {
		if accounts[account.Identity] {
			return fmt.Errorf("duplicate account identity %d", account.Identity)
		}

		accounts[account.Identity] = true
	}

	for _, transfer := range state.Transfers {
		if err := transfer.Validate(); err != nil {
			return err
		}
	}

	transfers := make(map[uint64]bool)
	for _, transfer := range state.Transfers {
		if transfers[transfer.Identity] {
			return fmt.Errorf("duplicate transfer identity %d", transfer.Identity)
		}

		transfers[transfer.Identity] = true
	}

	return nil
}
