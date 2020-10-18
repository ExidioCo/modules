package types

type GenesisState struct {
	Accounts  Accounts  `json:"accounts"`
	Transfers Transfers `json:"transfers"`
}

func NewGenesisState(accounts Accounts, transfers Transfers) GenesisState {
	return GenesisState{
		Accounts:  accounts,
		Transfers: transfers,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
