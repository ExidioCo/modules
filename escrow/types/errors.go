package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal        = errors.Register(ModuleName, 101, "failed to marshal")
	ErrorUnmarshal      = errors.Register(ModuleName, 102, "failed to unmarshal")
	ErrorUnknownMessage = errors.Register(ModuleName, 103, "unknown message")
	ErrorUnknownQuery   = errors.Register(ModuleName, 104, "unknown query")

	ErrorAccountDoesNotExist = errors.Register(ModuleName, 105, "account does not exist")
	ErrorEscrowDoesNotExist  = errors.Register(ModuleName, 106, "escrow does not exist")
	ErrorDeadlineExceeded    = errors.Register(ModuleName, 107, "deadline exceeded")
	ErrorEscrowFulfilled     = errors.Register(ModuleName, 108, "escrow fulfilled")
	ErrorDuplicateSigner     = errors.Register(ModuleName, 108, "duplicate signer")
	ErrorHolderDoesNotExist  = errors.Register(ModuleName, 109, "holder does not exist")
)
