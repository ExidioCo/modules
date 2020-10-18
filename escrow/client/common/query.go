package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/exidioco/modules/escrow/types"
)

func QueryEscrow(ctx context.CLIContext, identity uint64) (*types.Escrow, error) {
	params := types.NewQueryEscrowParams(identity)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryEscrow)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no escrow found")
	}

	var escrow types.Escrow
	if err := ctx.Codec.UnmarshalJSON(res, &escrow); err != nil {
		return nil, err
	}

	return &escrow, nil
}

func QueryEscrows(ctx context.CLIContext, page, limit int) (types.Escrows, error) {
	params := types.NewQueryEscrowsParams(page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryEscrows)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no escrows found")
	}

	var escrows types.Escrows
	if err := ctx.Codec.UnmarshalJSON(res, &escrows); err != nil {
		return nil, err
	}

	return escrows, nil
}
