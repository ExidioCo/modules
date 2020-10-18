package types

import (
	"time"

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
	AccountsCountKey = []byte{0x00}
	AccountKeyPrefix = []byte{0x01}

	TransfersCountKey            = []byte{0x10}
	TransferKeyPrefix            = []byte{0x11}
	TransferForDeadlineKeyPrefix = []byte{0x12}
)

func AccountKey(i uint64) []byte {
	return append(AccountKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func TransferKey(i uint64) []byte {
	return append(TransferKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func GetTransferForDeadlineKeyPrefix(at time.Time) []byte {
	return append(TransferForDeadlineKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func TransferForDeadlineKey(at time.Time, i uint64) []byte {
	return append(GetTransferForDeadlineKeyPrefix(at), sdk.Uint64ToBigEndian(i)...)
}
