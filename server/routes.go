package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Eldius/auth-server-go/auth"
	"github.com/Eldius/auth-server-go/logger"
	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/webapp-healthcheck-go/health"
)

/*
Start starts server
*/
func Start(appPort int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/login", auth.HandleLogin())
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Health check
	mux.HandleFunc("/health", health.BuildChecker([]health.ServiceChecker{
		health.NewDBChecker("main-db", repository.GetDB().DB(), time.Duration(2*time.Second)),
	}, map[string]string{
		"app": "auth-server-go",
	},
	))

	host := fmt.Sprintf(":%d", appPort)
	logger.Logger().Infof("starting app at '%s'", host)
	if err := http.ListenAndServe(host, mux); err != nil {
		logger.Logger().Panic(err.Error())
	}
}
