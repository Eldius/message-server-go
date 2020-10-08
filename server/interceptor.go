package server

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

func RequestIdInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())
		r = r.WithContext(ctx)
		log.Debugf("Incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, id.String())
		next.ServeHTTP(w, r)
		log.Debugf("Finished handling http req. %s", id.String())
	})
}