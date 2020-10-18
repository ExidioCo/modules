package joint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	k.IterateTransfersForDeadline(ctx, ctx.BlockTime(), func(_ int, item types.Transfer) bool {
		k.DeleteTransferForDeadline(ctx, item.Deadline, item.Identity)
		if err := k.Add(ctx, item.From, item.Coins); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
