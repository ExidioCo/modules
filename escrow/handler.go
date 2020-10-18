package escrow

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreate:
			return handleCreateEscrow(ctx, k, msg)
		case types.MsgApprove:
			return handleApprove(ctx, k, msg)
		default:
			return nil, types.ErrorUnknownMessage
		}
	}
}

func handleCreateEscrow(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreate) (*sdk.Result, error) {
	account, found := k.GetAccount(ctx, msg.From)
	if !found {
		return nil, types.ErrorAccountDoesNotExist
	}

	if err := k.Lock(ctx, account.Identity, msg.Coins); err != nil {
		return nil, err
	}

	count := k.GetCount(ctx)
	escrow := types.Escrow{
		Identity: count + 1,
		From:     msg.From,
		To:       msg.To,
		Coins:    msg.Coins,
		Deadline: msg.Deadline,
		Signers:  nil,
	}

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	k.SetEscrow(ctx, escrow)
	k.SetEscrowForDeadline(ctx, escrow.Deadline, escrow.Identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", escrow.Identity)),
		sdk.NewAttribute(types.AttributeKeyFrom, fmt.Sprintf("%d", escrow.From)),
		sdk.NewAttribute(types.AttributeKeyTo, escrow.To.String()),
		sdk.NewAttribute(types.AttributeKeyCoins, escrow.Coins.String()),
		sdk.NewAttribute(types.AttributeKeyDeadline, escrow.Deadline.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleApprove(ctx sdk.Context, k keeper.Keeper, msg types.MsgApprove) (*sdk.Result, error) {
	escrow, found := k.GetEscrow(ctx, msg.Identity)
	if !found {
		return nil, types.ErrorEscrowDoesNotExist
	}
	if ctx.BlockTime().After(escrow.Deadline) {
		return nil, types.ErrorDeadlineExceeded
	}

	for _, address := range escrow.Signers {
		if msg.Signer.Equals(address) {
			return nil, types.ErrorDuplicateSigner
		}
	}

	account, found := k.GetAccount(ctx, escrow.From)
	if !found {
		return nil, types.ErrorAccountDoesNotExist
	}

	if len(escrow.Signers) == account.Consents {
		return nil, types.ErrorEscrowFulfilled
	}

	index := 0
	for ; index < len(account.Holders); index++ {
		if msg.Signer.Equals(account.Holders[index]) {
			break
		}
	}

	if index == len(account.Holders) {
		return nil, types.ErrorHolderDoesNotExist
	}

	escrow.Signers = append(escrow.Signers, msg.Signer)

	k.SetEscrow(ctx, escrow)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeApprove,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", escrow.Identity)),
		sdk.NewAttribute(types.AttributeKeyHolder, msg.Signer.String()),
	))

	if len(escrow.Signers) == account.Consents {
		if err := k.SendCoinsToAccount(ctx, escrow.To, escrow.Coins); err != nil {
			return nil, err
		}

		k.DeleteEscrowForDeadline(ctx, escrow.Deadline, escrow.Identity)
	}

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
