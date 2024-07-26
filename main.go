package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/mcclean.xyz/reload"
	"github.com/ljmcclean/mcclean.xyz/views"
)

const port = ":1234"
const dev = true

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(views.Index("Eli")))
	mux.Handle("/palette", templ.Handler(views.Palette()))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	srv := &http.Server{
		Handler: mux,
	}

	if dev {
		mux.Handle("/reload", http.HandlerFunc(reload.Handler))
		srv.Handler = reload.BrowserReloadMiddleware(mux)
	}

	fmt.Println("Listening on port ", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	srv.Serve(ln)
}
