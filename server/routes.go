package server

import (
	"net/http"
	"time"

	"github.com/Eldius/message-server-go/auth"
	"github.com/Eldius/message-server-go/config"
	"github.com/Eldius/message-server-go/repository"
	"github.com/Eldius/webapp-healthcheck-go/health"
)

/*
Routes creates the app router
*/
func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/login", auth.HandleLogin())

	// Health check
	mux.HandleFunc("/health", health.BuildChecker([]health.ServiceChecker{
		health.NewDBChecker("main-db", repository.GetDB().DB(), time.Duration(2*time.Second)),
	}, map[string]string{
		"app": "message-server-go",
		"version": config.GetVersion(),
		"build-date": config.GetBuildDate(),
		"branch-name": config.GetBranchName(),
	}))

	mux.Handle("/admin", auth.AuthInterceptor(AdminHandler))
	mux.Handle("/message", auth.AuthInterceptor(MessageHandler))

	return mux
}
