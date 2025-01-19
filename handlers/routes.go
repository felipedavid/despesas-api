package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/middleware"
)

func SetupMultiplexer() http.Handler {
	mux := http.NewServeMux()

	HandleCustomHandler(mux, "/healthcheck", healthcheck)
	HandleCustomHandler(mux, "/user", handleRegisterUser)
	HandleCustomHandler(mux, "/auth/{provider}/callback", handleOauthCallback)
	HandleCustomHandler(mux, "/auth/{provider}", handleOauthAuthentication)

	return middleware.LogRequest(mux)
}

func HandleCustomHandler(mux *http.ServeMux, path string, h customHandler) {
	mux.HandleFunc(path, handleErrors(h))
}
