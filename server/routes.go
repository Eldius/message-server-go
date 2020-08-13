package server

import (
	"fmt"
	"net/http"
)

/*
Start starts server
*/
func Start(appPort int) {
	http.HandleFunc("/", IndexHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
}
