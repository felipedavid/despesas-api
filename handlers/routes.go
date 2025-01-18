package handlers

import "net/http"

func SetupMultiplexer() *http.ServeMux {
	mux := http.NewServeMux()

	HandleCustomHandler(mux, "/healthcheck", healthcheck)

	return mux
}

func HandleCustomHandler(mux *http.ServeMux, path string, h customHandler) {
	mux.HandleFunc(path, handleErrors(h))
}
