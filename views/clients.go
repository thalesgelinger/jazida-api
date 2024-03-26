package views

import (
	"jazida-api/db"
	"jazida-api/entities"
	"net/http"
)

type ClientProps struct {
	Clients []entities.Client
}

func Config(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	c, err := db.GetClients()
	if err != nil {
		// TODO: handle this later
		return
	}

	clients := ClientProps{
		Clients: c,
	}

	t.Render(w, "clients.html", clients)
}

func NewFormClient(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	t.Render(w, "new-client", "")
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	name := r.FormValue("name")
	plate := r.FormValue("plate")

	client := entities.Client{
		Name:  name,
		Plate: plate,
	}

	db.SaveClient(&client)

	c, err := db.GetClients()
	if err != nil {
		// TODO: handle this later
		return
	}

	clients := ClientProps{
		Clients: c,
	}

	t.Render(w, "clients", clients)
}
