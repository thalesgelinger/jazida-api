package views

import (
	"html/template"
	"io"
)

type Templates struct {
	template *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) {
	t.template.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		template: template.Must(template.ParseGlob("templates/*.html")),
	}
}
