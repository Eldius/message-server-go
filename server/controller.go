package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Eldius/jwt-auth-go/auth"
	authRepo "github.com/Eldius/jwt-auth-go/repository"
	"github.com/Eldius/jwt-auth-go/user"
	"github.com/Eldius/message-server-go/logger"
	"github.com/Eldius/message-server-go/messenger"
	"github.com/Eldius/message-server-go/repository"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// IndexHandler is the handler for index path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "Index", nil); err != nil {
		logger.Logger().Println(err.Error())
		w.WriteHeader(500)
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Logger()
	if r.Method == http.MethodPost {
		var mr *messenger.NewMessageRequest
		if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
			log.WithError(err).Error("FailedToSaveMessage")
			return
		}
		from := auth.GetCurrentUser(r)
		to := authRepo.FindUser(mr.To)
		m := messenger.NewMessageWithMessage(from.ID, to.ID, mr.Message)

		w.WriteHeader(201)
		repository.SaveMessage(m)
	} else if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		u := auth.GetCurrentUser(r)
		var msgs []messenger.MessageResponse
		for _, m := range repository.FindMessageTo(u.ID) {
			msgs = append(msgs, messenger.MessageResponse{
				ID:          m.ID,
				Destination: u.Name,
				From:        parseMessageOrigin(m),
				Message:     m.Message,
				Sent:        m.Sent,
			})
		}

		if err := json.NewEncoder(w).Encode(msgs); err != nil {
			log.WithError(err).Error("FailedToFetchMessages")
			return
		}
	}
	w.WriteHeader(405)
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
func parseMessageOrigin(m messenger.Message) string {
	fromUsr := authRepo.FindUserByID(m.From)
	if fromUsr != nil {
		return fromUsr.Name
	}
	return ""
}
