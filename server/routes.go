package server

import (
	"fmt"
	"net/http"

	"github.com/Eldius/auth-server-go/logger"
)

/*
Start starts server
*/
func Start(appPort int) {
	http.HandleFunc("/", IndexHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	logger.Logger().Infof("starting app at http://localhost:%d", appPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil); err != nil {
		logger.Logger().Panic(err.Error())
	}
}
