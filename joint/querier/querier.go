package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAccount:
			return queryAccount(ctx, req, k)
		case types.QueryAccounts:
			return queryAccounts(ctx, req, k)
		case types.QueryTransfer:
			return queryTransfer(ctx, req, k)
		case types.QueryTransfers:
			return queryTransfers(ctx, req, k)
		default:
			return nil, types.ErrorUnknownQuery
		}
	}
}
