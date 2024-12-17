package templates

import (
	"embed"
	"html/template"
)

var (
	//go:embed *.html
	templateFS embed.FS

	Templates *template.Template
)

func init() {
	Templates = template.Must(
		template.ParseFS(
			templateFS,
			"layout.html",
			"*.html",
		),
	)
}
