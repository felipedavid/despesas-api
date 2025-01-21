package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/middleware"
)

func SetupMultiplexer() http.Handler {
	mux := http.NewServeMux()

	routes := map[string]customHandler{
		"GET /healthcheck":                    healthcheck,
		"GET /auth/{provider}/callback":       handleOauthCallback,
		"GET /auth/{provider}":                handleOauthAuthentication,
		"POST /user":                          handleRegisterUser,
		"GET /transaction":                    handleListUserTransactions,
		"POST /transaction":                   handleCreateTransaction,
		"DELETE /transaction/{transactionID}": handleDeleteTransaction,
		"POST /account":                       handleCreateAccount,
		"DELETE /account/{accountID}":         handleDeleteAccount,
		"GET /account":                        handleListUserAccounts,
	}

	for path, handler := range routes {
		mux.HandleFunc(path, handleErrors(handler))
	}

	return middleware.LogRequest(middleware.SpecifyLanguage(mux))
}
