package routes

import (
	"html/template"
	"net/http"
)

func init() {}

func NewIndexHandlerX(t *template.Template) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		t.Execute(wr, nil)
	}
}

func NewIndexHandler(wr http.ResponseWriter, r *http.Request) {
	http.Redirect(wr, r, "/contacts/", http.StatusFound)
}
