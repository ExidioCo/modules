package escrow

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	var count uint64
	for _, escrow := range state {
		if escrow.Identity > count {
			count = escrow.Identity
		}

		k.SetEscrow(ctx, escrow)
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetEscrows(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, escrow := range state {
		if err := escrow.Validate(); err != nil {
			return err
		}
	}

	escrows := make(map[uint64]bool)
	for _, escrow := range state {
		if escrows[escrow.Identity] {
			return fmt.Errorf("duplicate escrow identity %d", escrow.Identity)
		}

		escrows[escrow.Identity] = true
	}

	return nil
}
