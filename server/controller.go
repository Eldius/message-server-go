package server

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// IndexHandler is the handler for index path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
