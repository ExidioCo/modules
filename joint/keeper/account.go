package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/exidioco/modules/joint/types"
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

func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	key := types.AccountKey(account.Identity)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(account)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetAccount(ctx sdk.Context, i uint64) (account types.Account, found bool) {
	store := ctx.KVStore(k.key)

	key := types.AccountKey(i)
	value := store.Get(key)
	if value == nil {
		return account, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &account)
	return account, true
}

func (k Keeper) GetAccounts(ctx sdk.Context) (items types.Accounts) {
	store := ctx.KVStore(k.key)

	iter := sdk.KVStorePrefixIterator(store, types.AccountKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Account
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) Add(ctx sdk.Context, i uint64, coins sdk.Coins) error {
	account, found := k.GetAccount(ctx, i)
	if !found {
		return types.ErrorAccountDoesNotExist
	}

	account.Coins = account.Coins.Add(coins...)
	k.SetAccount(ctx, account)

	return nil
}

func (k Keeper) Subtract(ctx sdk.Context, i uint64, coins sdk.Coins) error {
	account, found := k.GetAccount(ctx, i)
	if !found {
		return types.ErrorAccountDoesNotExist
	}

	account.Coins, _ = account.Coins.SafeSub(coins)
	if account.Coins.IsAnyNegative() {
		return types.ErrorNegativeCoins
	}

	k.SetAccount(ctx, account)
	return nil
}

func (k Keeper) Deposit(ctx sdk.Context, from sdk.AccAddress, to uint64, coins sdk.Coins) error {
	if err := k.supply.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins); err != nil {
		return err
	}

	return k.Add(ctx, to, coins)
}

func (k Keeper) Withdraw(ctx sdk.Context, from uint64, to sdk.AccAddress, coins sdk.Coins) error {
	if err := k.Subtract(ctx, from, coins); err != nil {
		return err
	}

	return k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
}

func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to uint64, coins sdk.Coins) error {
	if err := k.supply.SendCoinsFromModuleToModule(ctx, from, types.ModuleName, coins); err != nil {
		return err
	}

	return k.Add(ctx, to, coins)
}

func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, from uint64, to string, coins sdk.Coins) error {
	if err := k.Subtract(ctx, from, coins); err != nil {
		return err
	}

	return k.supply.SendCoinsFromModuleToModule(ctx, types.ModuleName, to, coins)
}
