package server

import (
	"context"
	"net/http"

	"github.com/Eldius/message-server-go/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Logger()
		origin := r.Header.Get("Origin")
		log.WithFields(logrus.Fields{
			"origin": r.Header.Get("Origin"),
			"dest":   r.URL.String(),
			"method": r.Method,
		}).Info("CORS")

		if origin != "" {
			log.Println("setupHeaders")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin, Host")
			w.WriteHeader(200)
		}

		if r.Method != http.MethodOptions {
			log.Println("executeRequest")
			next.ServeHTTP(w, r)
		}

	})
}
