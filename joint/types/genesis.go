package types

type GenesisState = Accounts

func NewGenesisState(accounts Accounts) GenesisState {
	return accounts
}

func DefaultGenesisState() GenesisState {
	return nil
}
