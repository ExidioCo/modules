package escrow

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	k.IterateEscrowsForDeadline(ctx, ctx.BlockTime(), func(_ int, item types.Escrow) bool {
		k.DeleteEscrowForDeadline(ctx, item.Deadline, item.Identity)
		if err := k.Unlock(ctx, item.From, item.Coins); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
