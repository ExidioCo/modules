package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/exidioco/modules/joint/types"
)

func QueryAccount(ctx context.CLIContext, identity uint64) (*types.Account, error) {
	params := types.NewQueryAccountParams(identity)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAccount)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no account found")
	}

	var account types.Account
	if err := ctx.Codec.UnmarshalJSON(res, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func QueryAccounts(ctx context.CLIContext, page, limit int) (types.Accounts, error) {
	params := types.NewQueryAccountsParams(page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAccounts)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no accounts found")
	}

	var accounts types.Accounts
	if err := ctx.Codec.UnmarshalJSON(res, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func QueryTransfer(ctx context.CLIContext, identity uint64) (*types.Transfer, error) {
	params := types.NewQueryTransferParams(identity)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTransfer)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no transfer found")
	}

	var transfer types.Transfer
	if err := ctx.Codec.UnmarshalJSON(res, &transfer); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func QueryTransfers(ctx context.CLIContext, page, limit int) (types.Transfers, error) {
	params := types.NewQueryTransfersParams(page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTransfers)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no transfers found")
	}

	var transfers types.Transfers
	if err := ctx.Codec.UnmarshalJSON(res, &transfers); err != nil {
		return nil, err
	}

	return transfers, nil
}
