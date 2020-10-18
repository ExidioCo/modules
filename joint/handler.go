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
		case types.MsgSend:
			return handleSend(ctx, k, msg)
		case types.MsgApprove:
			return handleApprove(ctx, k, msg)
		default:
			return nil, types.ErrorUnknownMessage
		}
	}
}

func handleCreate(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreate) (*sdk.Result, error) {
	count := k.GetAccountsCount(ctx)
	account := types.Account{
		Identity: count + 1,
		Consents: msg.Consents,
		Holders:  msg.Holders,
	}

	k.SetAccount(ctx, account)
	k.SetAccountsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreate,
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

func handleSend(ctx sdk.Context, k keeper.Keeper, msg types.MsgSend) (*sdk.Result, error) {
	account, found := k.GetAccount(ctx, msg.From)
	if !found {
		return nil, types.ErrorAccountDoesNotExist
	}

	if err := k.Subtract(ctx, account.Identity, msg.Coins); err != nil {
		return nil, err
	}

	count := k.GetTransfersCount(ctx)
	transfer := types.Transfer{
		Identity: count + 1,
		From:     msg.From,
		To:       msg.To,
		Coins:    msg.Coins,
		Deadline: msg.Deadline,
		Signers:  nil,
	}

	k.SetTransfer(ctx, transfer)
	k.SetTransfersCount(ctx, count+1)
	k.SetTransferForDeadline(ctx, transfer.Deadline, transfer.Identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSend,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", transfer.Identity)),
		sdk.NewAttribute(types.AttributeKeyFrom, fmt.Sprintf("%d", transfer.From)),
		sdk.NewAttribute(types.AttributeKeyTo, transfer.To.String()),
		sdk.NewAttribute(types.AttributeKeyCoins, transfer.Coins.String()),
		sdk.NewAttribute(types.AttributeKeyDeadline, transfer.Deadline.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleApprove(ctx sdk.Context, k keeper.Keeper, msg types.MsgApprove) (*sdk.Result, error) {
	transfer, found := k.GetTransfer(ctx, msg.Identity)
	if !found {
		return nil, types.ErrorTransferDoesNotExist
	}
	if ctx.BlockTime().After(transfer.Deadline) {
		return nil, types.ErrorDeadlineExceeded
	}

	for _, address := range transfer.Signers {
		if msg.Signer.Equals(address) {
			return nil, types.ErrorDuplicateSigner
		}
	}

	account, found := k.GetAccount(ctx, transfer.From)
	if !found {
		return nil, types.ErrorAccountDoesNotExist
	}

	if len(transfer.Signers) == account.Consents {
		return nil, types.ErrorTransferFulfilled
	}

	var index int
	for ; index < len(account.Holders); index++ {
		if msg.Signer.Equals(account.Holders[index]) {
			break
		}
	}

	if index == len(account.Holders) {
		return nil, types.ErrorHolderDoesNotExist
	}

	transfer.Signers = append(transfer.Signers, msg.Signer)

	k.SetTransfer(ctx, transfer)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeApprove,
		sdk.NewAttribute(types.AttributeKeyIdentity, fmt.Sprintf("%d", transfer.Identity)),
		sdk.NewAttribute(types.AttributeKeyHolder, msg.Signer.String()),
	))

	if len(transfer.Signers) == account.Consents {
		if err := k.Send(ctx, transfer.To, transfer.Coins); err != nil {
			return nil, err
		}

		k.DeleteTransferForDeadline(ctx, transfer.Deadline, transfer.Identity)
	}

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
