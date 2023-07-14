package config

import (
	"embed"
	"html/template"
	"io"
)

//go:embed view/*.html
var tmplFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	templates := template.Must(
		template.
			New("").
			Funcs(template.FuncMap{}).
			ParseFS(tmplFS, "view/*.html"))

	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "view/"+name))

	return tmpl.ExecuteTemplate(w, name, data)
}
