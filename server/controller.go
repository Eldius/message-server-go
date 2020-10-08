package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Eldius/auth-server-go/auth"
	"github.com/Eldius/auth-server-go/logger"
	"github.com/Eldius/auth-server-go/user"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// IndexHandler is the handler for index path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "Index", nil); err != nil {
		logger.Logger().Println(err.Error())
		w.WriteHeader(500)
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
