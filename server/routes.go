package server

import (
	"net/http"
	"time"

	"github.com/Eldius/jwt-auth-go/auth"
	authRep "github.com/Eldius/jwt-auth-go/repository"
	"github.com/Eldius/message-server-go/config"
	"github.com/Eldius/message-server-go/repository"
	"github.com/Eldius/webapp-healthcheck-go/health"
)

/*
Routes creates the app router
*/
func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.HandleLogin())
	mux.Handle("/user", auth.AuthInterceptor(auth.HandleNewUser()))

	// Health check
	mux.HandleFunc("/health", health.BuildChecker([]health.ServiceChecker{
		health.NewDBChecker("main-db", repository.GetDB().DB(), time.Duration(2*time.Second)),
		health.NewDBChecker("auth-db", authRep.GetDB().DB(), time.Duration(2*time.Second)),
	}, map[string]string{
		"app":         "message-server-go",
		"version":     config.GetVersion(),
		"build-date":  config.GetBuildDate(),
		"branch-name": config.GetBranchName(),
	}))

	mux.Handle("/admin", auth.AuthInterceptor(AdminHandler))
	mux.Handle("/message", auth.AuthInterceptor(MessageHandler))

	return mux
}
