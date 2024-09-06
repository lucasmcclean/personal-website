package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func NewServer(address string) *http.Server {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.Handle("/{$}", templ.Handler(Index()))

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server
}
