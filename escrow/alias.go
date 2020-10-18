// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/exidioco/modules/escrow/types
// ALIASGEN: github.com/exidioco/modules/escrow/keeper
// ALIASGEN: github.com/exidioco/modules/escrow/querier
package escrow

import (
	"github.com/exidioco/modules/escrow/keeper"
	"github.com/exidioco/modules/escrow/querier"
	"github.com/exidioco/modules/escrow/types"
)

const (
	EventTypeSetCount    = types.EventTypeSetCount
	EventTypeSet         = types.EventTypeSet
	EventTypeApprove     = types.EventTypeApprove
	AttributeKeyCount    = types.AttributeKeyCount
	AttributeKeyIdentity = types.AttributeKeyIdentity
	AttributeKeyFrom     = types.AttributeKeyFrom
	AttributeKeyTo       = types.AttributeKeyTo
	AttributeKeyCoins    = types.AttributeKeyCoins
	AttributeKeyDeadline = types.AttributeKeyDeadline
	AttributeKeyHolder   = types.AttributeKeyHolder
	ModuleName           = types.ModuleName
	QuerierRoute         = types.QuerierRoute
	RouterKey            = types.RouterKey
	StoreKey             = types.StoreKey
	QueryEscrow          = types.QueryEscrow
	QueryEscrows         = types.QueryEscrows
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	EscrowKey                     = types.EscrowKey
	GetEscrowForDeadlineKeyPrefix = types.GetEscrowForDeadlineKeyPrefix
	EscrowForDeadlineKey          = types.EscrowForDeadlineKey
	NewMsgCreate                  = types.NewMsgCreate
	NewMsgApprove                 = types.NewMsgApprove
	NewQueryEscrowParams          = types.NewQueryEscrowParams
	NewQueryEscrowsParams         = types.NewQueryEscrowsParams
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = querier.NewQuerier

	// variable aliases
	ModuleCdc                  = types.ModuleCdc
	ErrorMarshal               = types.ErrorMarshal
	ErrorUnmarshal             = types.ErrorUnmarshal
	ErrorUnknownMessage        = types.ErrorUnknownMessage
	ErrorUnknownQuery          = types.ErrorUnknownQuery
	ErrorAccountDoesNotExist   = types.ErrorAccountDoesNotExist
	ErrorEscrowDoesNotExist    = types.ErrorEscrowDoesNotExist
	ErrorDeadlineExceeded      = types.ErrorDeadlineExceeded
	ErrorEscrowFulfilled       = types.ErrorEscrowFulfilled
	ErrorDuplicateSigner       = types.ErrorDuplicateSigner
	ErrorHolderDoesNotExist    = types.ErrorHolderDoesNotExist
	EventModuleName            = types.EventModuleName
	CountKey                   = types.CountKey
	EscrowKeyPrefix            = types.EscrowKeyPrefix
	EscrowForDeadlineKeyPrefix = types.EscrowForDeadlineKeyPrefix
)

type (
	Escrow             = types.Escrow
	Escrows            = types.Escrows
	GenesisState       = types.GenesisState
	MsgCreate          = types.MsgCreate
	MsgApprove         = types.MsgApprove
	QueryEscrowParams  = types.QueryEscrowParams
	QueryEscrowsParams = types.QueryEscrowsParams
	Keeper             = keeper.Keeper
)
