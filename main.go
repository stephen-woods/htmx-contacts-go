package main

import (
	"fmt"
	"htmx-contacts/internal/repository"
	"htmx-contacts/internal/routes"
	"htmx-contacts/internal/templates"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world... contacts")

	contactRepo := repository.NewContactRepository()

	r := mux.NewRouter()
	r.HandleFunc("/contacts/new", routes.HandleGetContactsNew(templates.GetContactsNew, contactRepo)).Methods(http.MethodGet)
	r.HandleFunc("/contacts", routes.HandleGetContacts(templates.GetContacts, contactRepo)).Methods(http.MethodGet)
	r.HandleFunc("/", routes.NewIndexHandler).Methods(http.MethodGet)

	log.Println("App running on port: 3000")
	http.ListenAndServe(":3000", r)
}
