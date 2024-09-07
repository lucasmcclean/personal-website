package main

import (
	"net/http"
)

func NewServer(address string) *http.Server {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.Handle("/{$}", handleIndex())

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server
}

func handleIndex() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "assets/html/index.html")
		})
}
