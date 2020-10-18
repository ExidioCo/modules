package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Escrow struct {
	Identity uint64 `json:"identity"`

	From  uint64         `json:"from"`
	To    sdk.AccAddress `json:"to"`
	Coins sdk.Coins      `json:"coins"`

	Deadline time.Time        `json:"deadline"`
	Signers  []sdk.AccAddress `json:"signers"`
}

func (e Escrow) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Identity: %d
From:     %s
To:       %d
Coins:    %s
Deadline: %s
Signers:  %s
`), e.Identity, e.From, e.To, e.Coins, e.Deadline, e.Signers)
}

func (e Escrow) Validate() error {
	if e.Identity == 0 {
		return fmt.Errorf("identity is zero")
	}
	if e.From == 0 {
		return fmt.Errorf("from is zero")
	}
	if e.To == nil || e.To.Empty() {
		return fmt.Errorf("to is nil or empty")
	}
	if e.Deadline.IsZero() {
		return fmt.Errorf("deadline is zero")
	}
	if e.Signers == nil {
		return fmt.Errorf("signers is nil")
	}

	return nil
}

type Escrows []Escrow
