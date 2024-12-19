package routes

import (
	"htmx-contacts/internal/model"
	"htmx-contacts/internal/repository"
	"htmx-contacts/internal/templates"
	"log"
	"net/http"
)

type ContactPageInfo struct {
	Q        string
	Contacts []model.Contact
}

type ContactNewPageInfo struct {
	Contact model.Contact
}

func HandleGetContacts(repo repository.ContactRepository) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		var info ContactPageInfo

		info.Q = r.URL.Query().Get("q")

		if len(info.Q) > 0 {
			info.Contacts = repo.Search(info.Q)
		} else {
			info.Contacts = repo.All()
		}

		err := templates.GetContacts.ExecuteTemplate(wr, "contacts.html", info)
		if err != nil {
			log.Println(err)
			http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func HandleGetContactsNew(repo repository.ContactRepository) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		var info ContactNewPageInfo

		err := templates.GetContactsNew.ExecuteTemplate(wr, "contacts_new.html", info)
		if err != nil {
			log.Println(err)
			http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
