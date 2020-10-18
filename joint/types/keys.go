package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "joint"
	QuerierRoute = ModuleName
	RouterKey    = ModuleName
	StoreKey     = ModuleName
)

var (
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	CountKey         = []byte{0x00}
	AccountKeyPrefix = []byte{0x01}
)

func AccountKey(i uint64) []byte {
	return append(AccountKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}
