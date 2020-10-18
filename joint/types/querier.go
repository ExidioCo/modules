package types

const (
	QueryAccount   = "account"
	QueryAccounts  = "accounts"
	QueryTransfer  = "transfer"
	QueryTransfers = "transfers"
)

type QueryAccountParams struct {
	Identity uint64 `json:"identity"`
}

func NewQueryAccountParams(identity uint64) QueryAccountParams {
	return QueryAccountParams{
		Identity: identity,
	}
}

type QueryAccountsParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryAccountsParams(page, limit int) QueryAccountsParams {
	return QueryAccountsParams{
		Page:  page,
		Limit: limit,
	}
}

type QueryTransferParams struct {
	Identity uint64 `json:"identity"`
}

func NewQueryTransferParams(identity uint64) QueryTransferParams {
	return QueryTransferParams{
		Identity: identity,
	}
}

type QueryTransfersParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryTransfersParams(page, limit int) QueryTransfersParams {
	return QueryTransfersParams{
		Page:  page,
		Limit: limit,
	}
}
