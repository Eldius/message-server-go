package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Eldius/auth-server-go/logger"
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
	response := map[string]interface{}{
		"msg":   "response message!",
		"reqId": reqId,
	}
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(body)
}
