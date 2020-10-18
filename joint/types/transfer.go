package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Transfer struct {
	Identity uint64 `json:"identity"`

	From  uint64         `json:"from"`
	To    sdk.AccAddress `json:"to"`
	Coins sdk.Coins      `json:"coins"`

	Deadline time.Time        `json:"deadline"`
	Signers  []sdk.AccAddress `json:"signers"`
}

func (t Transfer) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Identity: %d
From:     %d
To:       %s
Coins:    %s
Deadline: %s
Signers:  %s
`), t.Identity, t.From, t.To, t.Coins, t.Deadline, t.Signers)
}

func (t Transfer) Validate() error {
	if t.Identity == 0 {
		return fmt.Errorf("identity is zero")
	}
	if t.From == 0 {
		return fmt.Errorf("from is zero")
	}
	if t.To == nil || t.To.Empty() {
		return fmt.Errorf("to is nil or empty")
	}
	if t.Deadline.IsZero() {
		return fmt.Errorf("deadline is zero")
	}
	if t.Signers == nil {
		return fmt.Errorf("signers is nil")
	}

	return nil
}

type Transfers []Transfer
