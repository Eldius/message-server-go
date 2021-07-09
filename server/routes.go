package server

import (
	"net/http"
	"time"

	"github.com/Eldius/webapp-healthcheck-go/health"
	"github.com/eldius/jwt-auth-go/auth"
	authRepo "github.com/eldius/jwt-auth-go/repository"
	"github.com/eldius/message-server-go/config"
	"github.com/eldius/message-server-go/logger"
	"github.com/eldius/message-server-go/repository"
)

/*
Routes creates the app router
*/
func Routes() http.Handler {
	log := logger.Logger()
	mux := http.NewServeMux()
	db := repository.GetDB()
	repo := authRepo.NewRepositoryCustom(db)
	svc := auth.NewAuthServiceCustom(repo)
	h := auth.NewAuthHandlerCustom(svc)
	mux.HandleFunc("/login", h.HandleLogin())
	mux.Handle("/user", h.AuthInterceptor(h.HandleNewUser()))

	// Health check
	_db, err := db.DB()
	if err != nil {
		log.WithError(err).
			Fatal("Routes")
	}
	mux.HandleFunc("/health", health.BuildChecker([]health.ServiceChecker{
		health.NewDBCheckerCustomQuery("main-db", _db, time.Duration(2*time.Second), "select 1"),
	}, map[string]string{
		"app":         "message-server-go",
		"version":     config.GetVersion(),
		"build-date":  config.GetBuildDate(),
		"branch-name": config.GetBranchName(),
	}))

	mux.Handle("/admin", h.AuthInterceptor(AdminHandler))
	mux.Handle("/message", h.AuthInterceptor(MessageHandler(svc, repo)))

	return mux
}
