package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal        = errors.Register(ModuleName, 101, "failed to marshal")
	ErrorUnmarshal      = errors.Register(ModuleName, 102, "failed to unmarshal")
	ErrorUnknownMessage = errors.Register(ModuleName, 103, "unknown message")
	ErrorUnknownQuery   = errors.Register(ModuleName, 104, "unknown query")

	ErrorNegativeCoins       = errors.Register(ModuleName, 105, "negative coins")
	ErrorAccountDoesNotExist = errors.Register(ModuleName, 106, "address does not exist")
)
