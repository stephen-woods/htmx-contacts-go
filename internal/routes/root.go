package routes

import (
	"html/template"
	"net/http"
)

func init() {}

func NewIndexHandler(t *template.Template) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		t.Execute(wr, nil)
	}
}
