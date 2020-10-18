package types

type GenesisState = Escrows

func NewGenesisState(escrows Escrows) GenesisState {
	return escrows
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
