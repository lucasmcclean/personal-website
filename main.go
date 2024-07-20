package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/mcclean.xyz/views"
)

const port = ":1234"

func main() {
	http.Handle("/", templ.Handler(views.Index("World")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Listening on port ", port)
	http.ListenAndServe(port, nil)
}
