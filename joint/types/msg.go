package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgCreate)(nil)
	_ sdk.Msg = (*MsgDeposit)(nil)
	_ sdk.Msg = (*MsgSend)(nil)
	_ sdk.Msg = (*MsgApprove)(nil)
)

type MsgCreate struct {
	Signer sdk.AccAddress `json:"signer"`

	Consents int              `json:"consents"`
	Holders  []sdk.AccAddress `json:"holders"`
}

func NewMsgCreate(signer sdk.AccAddress, consents int, holders []sdk.AccAddress) MsgCreate {
	return MsgCreate{
		Signer:   signer,
		Consents: consents,
		Holders:  holders,
	}
}

func (m MsgCreate) Route() string {
	return RouterKey
}

func (m MsgCreate) Type() string {
	return "create"
}

func (m MsgCreate) ValidateBasic() error {
	if m.Signer == nil || m.Signer.Empty() {
		return fmt.Errorf("signer is nil or empty")
	}
	if m.Consents == 0 {
		return fmt.Errorf("consents is zero")
	}
	if m.Holders == nil {
		return fmt.Errorf("holders is nil")
	}
	if m.Consents > len(m.Holders) {
		return fmt.Errorf("consents is greater than holders length")
	}

	return nil
}

func (m MsgCreate) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeposit struct {
	Signer sdk.AccAddress `json:"signer"`

	Identity uint64    `json:"identity"`
	Coins    sdk.Coins `json:"coins"`
}

func NewMsgDeposit(signer sdk.AccAddress, identity uint64, coins sdk.Coins) MsgDeposit {
	return MsgDeposit{
		Signer:   signer,
		Identity: identity,
		Coins:    coins,
	}
}

func (m MsgDeposit) Route() string {
	return RouterKey
}

func (m MsgDeposit) Type() string {
	return "deposit"
}

func (m MsgDeposit) ValidateBasic() error {
	if m.Signer == nil || m.Signer.Empty() {
		return fmt.Errorf("signer is nil or empty")
	}
	if m.Identity == 0 {
		return fmt.Errorf("identity is zero")
	}
	if m.Coins == nil || !m.Coins.IsValid() {
		return fmt.Errorf("coins is nil or invalid")
	}

	return nil
}

func (m MsgDeposit) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgSend struct {
	Signer sdk.AccAddress `json:"signer"`

	From     uint64         `json:"from"`
	To       sdk.AccAddress `json:"to"`
	Coins    sdk.Coins      `json:"coins"`
	Deadline time.Time      `json:"deadline"`
}

func NewMsgSend(signer sdk.AccAddress, from uint64, to sdk.AccAddress, coins sdk.Coins, deadline time.Time) MsgSend {
	return MsgSend{
		Signer:   signer,
		From:     from,
		To:       to,
		Coins:    coins,
		Deadline: deadline,
	}
}

func (m MsgSend) Route() string {
	return RouterKey
}

func (m MsgSend) Type() string {
	return "send"
}

func (m MsgSend) ValidateBasic() error {
	if m.Signer == nil || m.Signer.Empty() {
		return fmt.Errorf("signer is nil or empty")
	}
	if m.From == 0 {
		return fmt.Errorf("from is zero")
	}
	if m.To == nil || m.To.Empty() {
		return fmt.Errorf("to is nil or empty")
	}
	if m.Coins == nil || !m.Coins.IsValid() {
		return fmt.Errorf("coins is nil or empty")
	}
	if m.Deadline.IsZero() {
		return fmt.Errorf("deadline is zero")
	}

	return nil
}

func (m MsgSend) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgApprove struct {
	Signer sdk.AccAddress `json:"signer"`

	Identity uint64 `json:"identity"`
}

func NewMsgApprove(signer sdk.AccAddress, identity uint64) MsgApprove {
	return MsgApprove{
		Signer:   signer,
		Identity: identity,
	}
}

func (m MsgApprove) Route() string {
	return RouterKey
}

func (m MsgApprove) Type() string {
	return "approve"
}

func (m MsgApprove) ValidateBasic() error {
	if m.Signer == nil || m.Signer.Empty() {
		return fmt.Errorf("signer is nil or empty")
	}
	if m.Identity == 0 {
		return fmt.Errorf("identity is zero")
	}

	return nil
}

func (m MsgApprove) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgApprove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
