package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Eldius/auth-server-go/logger"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

/*
HandleLogin handles login requests
*/
func HandleLogin() http.HandlerFunc {
	log := logger.Logger()
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Type", "plain/text")
		var u LoginRequest
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.WithError(err).
				WithField("details", "Failed to parse request body").
				Println("HandleLogin")
			w.WriteHeader(401)
			return
		}
		if u.User == "" || u.Pass == "" {
			log.WithField("details", "empty user or pass").
				Println("HandleLogin")
			w.WriteHeader(401)
			return
		}
		cred, err := ValidatePass(u.User, u.Pass)
		if err != nil {
			log.WithError(err).
				WithFields(logrus.Fields{
					"details": "Failed to validate credentials",
					"user":    u.User,
				}).
				Println("HandleLogin")
			w.WriteHeader(401)
			return
		}
		log.Println(cred.User)
		if err != nil {
			log.WithError(err).
				WithFields(logrus.Fields{
					"details": "Failed to validate credentials",
					"user":    u.User,
				}).
				Println("HandleLogin")
			w.WriteHeader(401)
			return
		}

		token, err := ToJWT(*cred)
		if err != nil {
			log.WithError(err).Println("Failed to generate token")
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(token))
	}
}

func AuthInterceptor(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Logger()
		authHeader := r.Header.Get("Authorization")
		// TODO remove this before release
		if strings.HasPrefix(authHeader, "Bearer ") {
			log.WithField("header", authHeader).Println("authInterceptor")
			f.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	})
}
