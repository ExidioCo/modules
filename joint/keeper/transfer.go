package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/types"
)

func (k Keeper) SetTransfersCount(ctx sdk.Context, count uint64) {
	key := types.TransfersCountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetTransfersCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.key)

	key := types.TransfersCountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetTransfer(ctx sdk.Context, transfer types.Transfer) {
	key := types.TransferKey(transfer.Identity)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(transfer)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetTransfer(ctx sdk.Context, i uint64) (transfer types.Transfer, found bool) {
	store := ctx.KVStore(k.key)

	key := types.TransferKey(i)
	value := store.Get(key)
	if value == nil {
		return transfer, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &transfer)
	return transfer, true
}

func (k Keeper) GetTransfers(ctx sdk.Context) (items types.Transfers) {
	store := ctx.KVStore(k.key)

	iter := sdk.KVStorePrefixIterator(store, types.TransferKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Transfer
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetTransferForDeadline(ctx sdk.Context, at time.Time, i uint64) {
	key := types.TransferForDeadlineKey(at, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(i)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) DeleteTransferForDeadline(ctx sdk.Context, at time.Time, i uint64) {
	key := types.TransferForDeadlineKey(at, i)

	store := ctx.KVStore(k.key)
	store.Delete(key)
}

func (k Keeper) IterateTransfersForDeadline(ctx sdk.Context, end time.Time, fn func(index int, item types.Transfer) bool) {
	store := ctx.KVStore(k.key)

	iter := store.Iterator(types.TransferForDeadlineKeyPrefix, sdk.PrefixEndBytes(types.GetTransferForDeadlineKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var identity uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &identity)

		transfer, _ := k.GetTransfer(ctx, identity)
		if stop := fn(i, transfer); stop {
			break
		}
		i++
	}
}
