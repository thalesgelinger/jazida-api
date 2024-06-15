package views

import (
	"html/template"
	"io"
	"jazida-api/internal/infra/db"
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

type ViewHandler struct {
	db *db.Queries
}

func NewViewHandler(db *db.Queries) *ViewHandler {
	return &ViewHandler{db}
}
