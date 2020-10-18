package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func queryTransfer(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryTransferParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	transfer, found := k.GetTransfer(ctx, params.Identity)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(transfer)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}

func queryTransfers(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryTransfersParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	transfers := k.GetTransfers(ctx)

	start, end := client.Paginate(len(transfers), params.Page, params.Limit, len(transfers))
	if start < 0 || end < 0 {
		transfers = types.Transfers{}
	} else {
		transfers = transfers[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(transfers)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}
