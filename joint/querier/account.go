package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func queryAccount(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryAccountParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	account, found := k.GetAccount(ctx, params.Identity)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(account)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}

func queryAccounts(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryAccountsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal
	}

	accounts := k.GetAccounts(ctx)

	start, end := client.Paginate(len(accounts), params.Page, params.Limit, len(accounts))
	if start < 0 || end < 0 {
		accounts = types.Accounts{}
	} else {
		accounts = accounts[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(accounts)
	if err != nil {
		return nil, types.ErrorMarshal
	}

	return res, nil
}
