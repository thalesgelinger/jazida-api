package views

import (
	"net/http"
)

type Count struct {
	Count int
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()
	count := Count{1}
	t.Render(w, "home", count)
}
