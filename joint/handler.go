package joint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/keeper"
	"github.com/exidioco/modules/joint/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreate:
			return handleCreate(ctx, k, msg)
		case types.MsgDeposit:
			return handleDeposit(ctx, k, msg)
		default:
			return nil, types.ErrorUnknownMessage
		}
	}
}

func handleCreate(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreate) (*sdk.Result, error) {
	count := k.GetCount(ctx)
	account := types.Account{
		Identity: count + 1,
		Consents: msg.Consents,
		Holders:  msg.Holders,
	}

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	k.SetAccount(ctx, account)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", account.Identity)),
		sdk.NewAttribute(types.AttributeKeyConsents, fmt.Sprintf("%d", account.Consents)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleDeposit(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeposit) (*sdk.Result, error) {
	account, found := k.GetAccount(ctx, msg.Identity)
	if !found {
		return nil, types.ErrorAccountDoesNotExist
	}

	if err := k.Deposit(ctx, msg.Signer, account.Identity, msg.Coins); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDeposit,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", account.Identity)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
