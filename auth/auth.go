package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Eldius/auth-server-go/logger"
	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/auth-server-go/user"
	"github.com/sirupsen/logrus"
)

// ValidatePass validates user credentials
func ValidatePass(username string, pass string) (u *user.CredentialInfo, err error) {
	var usr = repository.FindUser(username)
	if usr.Hash == nil {
		err = fmt.Errorf("Failed to authenticate user")
		return
	}

	var ph []byte
	ph, err = user.Hash(pass, usr.Salt)
	if err != nil {
		return
	}

	if string(ph) == string(usr.Hash) {
		u = usr
	} else {
		err = fmt.Errorf("Failed to authenticate user")
	}

	return
}

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
		log.Println(cred)
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	}
}
