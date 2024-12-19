package main

import (
	"fmt"
	"htmx-contacts/internal/repository"
	"htmx-contacts/internal/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world... contacts")

	contactRepo := repository.NewContactRepository()

	r := http.NewServeMux()

	r.HandleFunc("GET /contacts/new", routes.HandleGetContactsNew(contactRepo))
	r.HandleFunc("GET /contacts", routes.HandleGetContacts(contactRepo))
	r.HandleFunc("GET /", routes.NewIndexHandler)

	log.Println("App running on port: 3000")
	http.ListenAndServe(":3000", r)
}
