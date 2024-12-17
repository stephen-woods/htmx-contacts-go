package routes

import (
	"html/template"
	"htmx-contacts/internal/model"
	"htmx-contacts/internal/repository"
	"net/http"
)

func NewContactsHandler(t *template.Template, repo repository.ContactRepository) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		var contacts []model.Contact
		text := r.URL.Query().Get("q")

		if len(text) > 0 {
			contacts = repo.Search(text)
		} else {
			contacts = repo.All()
		}

		err := t.ExecuteTemplate(wr, "contacts.html", contacts)
		if err != nil {
			http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
