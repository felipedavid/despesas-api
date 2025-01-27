package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/middleware"
)

func SetupMultiplexer() http.Handler {
	mux := http.NewServeMux()

	routes := map[string]customHandler{
		"GET /healthcheck":              healthcheck,
		"GET /auth/{provider}/callback": oauthCallback,
		"GET /auth/{provider}":          oauthAuthentication,
		"POST /auth":                    credentialsAuthentication,

		"POST /user": registerUser,

		"GET /transaction":                    listUserTransactions,
		"POST /transaction":                   createTransaction,
		"DELETE /transaction/{transactionID}": deleteTransaction,
		"PATCH /transaction/{transactionID}":  editTransaction,

		"GET /account":                listUserAccounts,
		"POST /account":               createAccount,
		"DELETE /account/{accountID}": deleteAccount,
		"PATCH /account/{accountID}":  editAccount,

		"POST /category":                createCategory,
		"GET /category":                 listUserCategories,
		"DELETE /category/{categoryID}": deleteCategory,
		"PATCH /category/{categoryID}":  editCategory,

		"GET /plan":        listAvailablePlans,
		"GET /active-plan": getActivePlan,
	}

	for path, handler := range routes {
		mux.HandleFunc(path, handleErrors(handler))
	}

	return middleware.LogRequest(middleware.Auth(middleware.SpecifyLanguage(mux)))
}
