package server

import (
	"net/http"
)

func New(address string) *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	var handler http.Handler = mux

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return server
}
