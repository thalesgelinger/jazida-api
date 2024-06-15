package views

import (
	"jazida-api/internal/infra/db"
	"log/slog"
	"net/http"
)

type HomeProps struct {
	Loads []db.Load
}

func (v *ViewHandler) Home(w http.ResponseWriter, r *http.Request) {
	slog.Info("Render home page")
	t := NewTemplate()

	allLoads, err := v.db.GetLoads(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("Error rendering home page", err)
		return
	}

	loads := allLoads

	homeProps := HomeProps{
		Loads: loads,
	}

	t.Render(w, "home", homeProps)
}
