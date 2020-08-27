package server

import (
	"fmt"
	"log"
	"net/http"
)

/*
Start starts server
*/
func Start(appPort int) {
	http.HandleFunc("/", IndexHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil); err != nil {
		log.Panic(err.Error())
	}
}
