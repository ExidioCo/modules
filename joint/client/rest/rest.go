package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/joint/accounts", queryAccounts(ctx)).
		Methods("GET")
	router.HandleFunc("/joint/accounts/{identity}", queryAccount(ctx)).
		Methods("GET")
}

func registerTxRoutes(_ context.CLIContext, _ *mux.Router) {
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}
