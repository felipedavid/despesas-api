package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/middleware"
)

func SetupMultiplexer() http.Handler {
	mux := http.NewServeMux()

	type HandlerSettings struct {
		handler       customHandler
		authenticated bool
	}

	routes := map[string]HandlerSettings{
		"GET /healthcheck":              {handler: healthcheck},
		"GET /auth/{provider}/callback": {handler: oauthCallback},
		"GET /auth/{provider}":          {handler: oauthAuthentication},
		"POST /auth":                    {handler: credentialsAuthentication},

		"POST /user": {handler: registerUser},

		"GET /transaction":                    {handler: listUserTransactions, authenticated: true},
		"GET /transaction/{transactionID}":    {handler: getUserTransaction, authenticated: true},
		"POST /transaction":                   {handler: createTransaction, authenticated: true},
		"DELETE /transaction/{transactionID}": {handler: deleteTransaction, authenticated: true},
		"PATCH /transaction/{transactionID}":  {handler: editTransaction, authenticated: true},

		"GET /account":                {handler: listUserAccounts, authenticated: true},
		"GET /account/{accountID}":    {handler: getUserAccount, authenticated: true},
		"POST /account":               {handler: createAccount, authenticated: true},
		"DELETE /account/{accountID}": {handler: deleteAccount, authenticated: true},
		"PATCH /account/{accountID}":  {handler: editAccount, authenticated: true},

		"POST /category":                {handler: createCategory, authenticated: true},
		"GET /category":                 {handler: listUserCategories, authenticated: true},
		"DELETE /category/{categoryID}": {handler: deleteCategory, authenticated: true},
		"PATCH /category/{categoryID}":  {handler: editCategory, authenticated: true},

		"GET /plan":        {handler: listAvailablePlans},
		"GET /active-plan": {handler: getActivePlan, authenticated: true},
	}

	for path, s := range routes {
		fn := handleErrors(s.handler)
		if s.authenticated {
			mux.Handle(path, middleware.RequireAuthentication(fn))
		} else {
			mux.HandleFunc(path, fn)
		}
	}

	return middleware.LogRequest(middleware.Auth(middleware.SpecifyLanguage(mux)))
}
