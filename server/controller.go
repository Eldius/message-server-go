package server

import (
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
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write([]byte("{\"msg\": \"response message!\"}"))
}
