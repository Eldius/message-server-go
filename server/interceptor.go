package server

import (
	"context"
	"net/http"

	"github.com/Eldius/auth-server-go/logger"
	"github.com/google/uuid"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

func RequestIdInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Logger()
		ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())
		r = r.WithContext(ctx)
		log.WithField("req", id.String()).Debugf("Incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, id.String())
		next.ServeHTTP(w, r)
		log.WithField("req", id.String()).Debugf("Finished handling http req. %s", id.String())
	})
}
