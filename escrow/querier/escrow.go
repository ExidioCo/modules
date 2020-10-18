package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/types"
)

func queryEscrow(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryEscrowParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	escrow, found := k.GetEscrow(ctx, params.Identity)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(escrow)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}

func queryEscrows(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryEscrowsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	escrows := k.GetEscrows(ctx)

	start, end := client.Paginate(len(escrows), params.Page, params.Limit, len(escrows))
	if start < 0 || end < 0 {
		escrows = types.Escrows{}
	} else {
		escrows = escrows[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(escrows)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}
