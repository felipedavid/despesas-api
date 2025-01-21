package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/middleware"
)

func SetupMultiplexer() http.Handler {
	mux := http.NewServeMux()

	routes := map[string]customHandler{
		"GET /healthcheck":                    healthcheck,
		"GET /auth/{provider}/callback":       oauthCallback,
		"GET /auth/{provider}":                oauthAuthentication,
		"POST /auth":                          credentialsAuthentication,
		"POST /user":                          registerUser,
		"GET /transaction":                    listUserTransactions,
		"POST /transaction":                   createTransaction,
		"DELETE /transaction/{transactionID}": deleteTransaction,
		"POST /account":                       createAccount,
		"DELETE /account/{accountID}":         deleteAccount,
		"GET /account":                        listUserAccounts,
	}

	for path, handler := range routes {
		mux.HandleFunc(path, handleErrors(handler))
	}

	return sessionManager.LoadAndSave(middleware.LogRequest(middleware.SpecifyLanguage(mux)))
}
