package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryEscrow:
			return queryEscrow(ctx, req, k)
		case types.QueryEscrows:
			return queryEscrows(ctx, req, k)
		default:
			return nil, types.ErrorUnknownQuery
		}
	}
}
