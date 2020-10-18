package types

const (
	QueryAccount  = "account"
	QueryAccounts = "accounts"
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
