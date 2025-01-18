package handlers

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	HandleCustomHandler(mux, "/healthcheck", healthcheck)

	return mux
}

func HandleCustomHandler(mux *http.ServeMux, path string, h customHandler) {
	mux.HandleFunc(path, httpHandler(h))
}
