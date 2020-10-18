package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "escrow"
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
	CountKey                   = []byte{0x00}
	EscrowKeyPrefix            = []byte{0x01}
	EscrowForDeadlineKeyPrefix = []byte{0x02}
)

func EscrowKey(i uint64) []byte {
	return append(EscrowKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func GetEscrowForDeadlineKeyPrefix(at time.Time) []byte {
	return append(EscrowForDeadlineKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func EscrowForDeadlineKey(at time.Time, i uint64) []byte {
	return append(GetEscrowForDeadlineKeyPrefix(at), sdk.Uint64ToBigEndian(i)...)
}
