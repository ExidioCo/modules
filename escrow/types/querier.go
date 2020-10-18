package types

const (
	QueryEscrow  = "escrow"
	QueryEscrows = "escrows"
)

type QueryEscrowParams struct {
	Identity uint64 `json:"identity"`
}

func NewQueryEscrowParams(identity uint64) QueryEscrowParams {
	return QueryEscrowParams{
		Identity: identity,
	}
}

type QueryEscrowsParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryEscrowsParams(page, limit int) QueryEscrowsParams {
	return QueryEscrowsParams{
		Page:  page,
		Limit: limit,
	}
}
