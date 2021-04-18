package server

import (
	"encoding/json"
	"net/http"

	"github.com/eldius/jwt-auth-go/auth"
	authRepo "github.com/eldius/jwt-auth-go/repository"
	"github.com/eldius/jwt-auth-go/user"
	"github.com/eldius/message-server-go/logger"
	"github.com/eldius/message-server-go/messenger"
	"github.com/eldius/message-server-go/repository"
)

func MessageHandler(svc *auth.AuthService, repo *authRepo.AuthRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.Logger()
		if r.Method == http.MethodPost {
			var mr *messenger.NewMessageRequest
			if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
				log.WithError(err).Error("FailedToSaveMessage")
				return
			}
			from := svc.GetCurrentUser(r)
			to := repo.FindUser(mr.To)
			m := messenger.NewMessageWithMessage(from.ID, to.ID, mr.Message)

			w.WriteHeader(201)
			repository.SaveMessage(m)
		} else if r.Method == http.MethodGet {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "application/json")
			u := svc.GetCurrentUser(r)
			var msgs []messenger.MessageResponse
			for _, m := range repository.FindMessageTo(u.ID) {
				msgs = append(msgs, messenger.MessageResponse{
					ID:          m.ID,
					Destination: u.Name,
					From:        parseMessageOrigin(m, repo),
					Message:     m.Message,
					Sent:        m.Sent,
				})
			}

			if err := json.NewEncoder(w).Encode(msgs); err != nil {
				log.WithError(err).Error("FailedToFetchMessages")
				return
			}
		}
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	reqId := r.Context().Value(ContextKeyRequestID)
	u := r.Context().Value(auth.CurrentUserKey).(*user.CredentialInfo)
	response := map[string]interface{}{
		"msg":   "response message!",
		"reqId": reqId,
		"user":  u.User,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// TODO think about use a cache solution here
func parseMessageOrigin(m messenger.Message, repo *authRepo.AuthRepository) string {
	fromUsr := repo.FindUserByID(m.From)
	if fromUsr != nil {
		return fromUsr.Name
	}
	return ""
}
