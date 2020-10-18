package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	Identity uint64           `json:"identity"`
	Consents int              `json:"consents"`
	Holders  []sdk.AccAddress `json:"holders"`
	Coins    sdk.Coins        `json:"coins"`
}

func (a Account) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Identity: %d
Consents: %d
Holders:  %s
Coins:    %s
`), a.Identity, a.Consents, a.Holders, a.Coins)
}

func (a Account) Validate() error {
	if a.Identity == 0 {
		return fmt.Errorf("identity is zero")
	}
	if a.Consents == 0 {
		return fmt.Errorf("consents is zero")
	}
	if a.Holders == nil {
		return fmt.Errorf("holders is nil")
	}
	if a.Consents > len(a.Holders) {
		return fmt.Errorf("consents is greater than holders length")
	}
	if a.Coins != nil && !a.Coins.IsValid() {
		return fmt.Errorf("coins are nil or invalid")
	}

	return nil
}

type Accounts []Account
