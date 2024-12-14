package main

import (
	"fmt"
	"htmx-contacts/internal/routes"
	"htmx-contacts/internal/templates"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world... contacts")

	templates := templates.Templates()
	http.Handle("GET /", routes.NewIndexHandler(templates["index"]))

	log.Println("App running on port: 3000")
	http.ListenAndServe(":3000", nil)
}
