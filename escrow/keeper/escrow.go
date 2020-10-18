package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/escrow/types"
)

func (k Keeper) SetCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.key)

	key := types.CountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetEscrow(ctx sdk.Context, escrow types.Escrow) {
	key := types.EscrowKey(escrow.Identity)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(escrow)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetEscrow(ctx sdk.Context, i uint64) (escrow types.Escrow, found bool) {
	store := ctx.KVStore(k.key)

	key := types.EscrowKey(i)
	value := store.Get(key)
	if value == nil {
		return escrow, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &escrow)
	return escrow, true
}

func (k Keeper) GetEscrows(ctx sdk.Context) (items types.Escrows) {
	store := ctx.KVStore(k.key)

	iter := sdk.KVStorePrefixIterator(store, types.EscrowKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Escrow
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetEscrowForDeadline(ctx sdk.Context, at time.Time, i uint64) {
	key := types.EscrowForDeadlineKey(at, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(i)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) DeleteEscrowForDeadline(ctx sdk.Context, at time.Time, i uint64) {
	key := types.EscrowForDeadlineKey(at, i)

	store := ctx.KVStore(k.key)
	store.Delete(key)
}

func (k Keeper) IterateEscrowsForDeadline(ctx sdk.Context, end time.Time, fn func(index int, item types.Escrow) bool) {
	store := ctx.KVStore(k.key)

	iter := store.Iterator(types.EscrowForDeadlineKeyPrefix, sdk.PrefixEndBytes(types.GetEscrowForDeadlineKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var identity uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &identity)

		escrow, _ := k.GetEscrow(ctx, identity)
		if stop := fn(i, escrow); stop {
			break
		}
		i++
	}
}
