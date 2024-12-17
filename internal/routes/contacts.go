package routes

import (
	"html/template"
	"htmx-contacts/internal/model"
	"htmx-contacts/internal/repository"
	"net/http"
)

type ContactPageInfo struct {
	Q        string
	Contacts []model.Contact
}

func NewContactsHandler(t *template.Template, repo repository.ContactRepository) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		var info ContactPageInfo

		info.Q = r.URL.Query().Get("q")

		if len(info.Q) > 0 {
			info.Contacts = repo.Search(info.Q)
		} else {
			info.Contacts = repo.All()
		}

		err := t.ExecuteTemplate(wr, "contacts.html", info)
		if err != nil {
			http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
