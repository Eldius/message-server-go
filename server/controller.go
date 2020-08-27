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
