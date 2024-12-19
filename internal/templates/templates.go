package templates

import (
	"embed"
	"html/template"
)

var (
	//go:embed *.html
	templateFS embed.FS

	GetContacts    *template.Template
	GetContactsNew *template.Template
)

func init() {
	GetContacts = template.Must(
		template.ParseFS(templateFS, "layout.html", "contacts.html"),
	)

	GetContactsNew = template.Must(
		template.ParseFS(templateFS, "layout.html", "contacts_new.html"),
	)
}
