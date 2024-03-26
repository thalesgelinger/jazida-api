package views

import (
	"jazida-api/db"
	"jazida-api/entities"
	"net/http"
)

type HomeProps struct {
	Loads []entities.Load
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	allLoads, err := db.GetLoads()
	if err != nil {
		// TODO: handle this later
		return
	}

	loads := allLoads

	homeProps := HomeProps{
		Loads: loads,
	}

	t.Render(w, "home", homeProps)
}

