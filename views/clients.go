package views

import (
	"jazida-api/internal/entity"
	"jazida-api/internal/infra/db"
	"net/http"
)

type ClientProps struct {
	Clients []entity.Client
}

func (v *ViewHandler) Config(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	clientsRow, err := v.db.GetClients(r.Context())
	if err != nil {
		// TODO: handle this later
		return
	}

	var clients []entity.Client

	for _, row := range clientsRow {
		clients = append(clients, entity.Client{
			Name:  row.Name,
			Plate: row.Plate,
		})
	}

	clientsProps := ClientProps{
		Clients: clients,
	}

	t.Render(w, "clients.html", clientsProps)
}

func (v *ViewHandler) NewFormClient(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	t.Render(w, "new-client", "")
}

func (v *ViewHandler) AddClient(w http.ResponseWriter, r *http.Request) {
	t := NewTemplate()

	name := r.FormValue("name")
	plate := r.FormValue("plate")

	client := db.AddClientParams{
		Name:  name,
		Plate: plate,
	}

	v.db.AddClient(r.Context(), client)

	clientsRow, err := v.db.GetClients(r.Context())
	if err != nil {
		// TODO: handle this later
		return
	}

	var clients []entity.Client

	for _, row := range clientsRow {
		clients = append(clients, entity.Client{
			Name:  row.Name,
			Plate: row.Plate,
		})
	}

	clientsProps := ClientProps{
		Clients: clients,
	}

	t.Render(w, "clients", clientsProps)
}
