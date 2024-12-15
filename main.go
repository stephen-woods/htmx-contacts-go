package main

import (
	"fmt"
	"htmx-contacts/internal/repository"
	"htmx-contacts/internal/routes"
	"htmx-contacts/internal/templates"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world... contacts")

	contactRepo := repository.NewContactRepository()

	templates := templates.Templates()
	_ = templates

	// http.Handle("GET /", routes.NewIndexHandler(templates["index"]))

	http.HandleFunc("/contacts", routes.NewContactsHandler(templates["contacts"], contactRepo))
	http.HandleFunc("/", routes.NewIndexHandler)

	log.Println("App running on port: 3000")
	http.ListenAndServe(":3000", nil)
}
