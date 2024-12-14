package templates

import "html/template"

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	templates["index"] = template.Must(template.ParseFiles("internal/templates/index.html"))
}

func Templates() map[string]*template.Template {
	return templates
}
