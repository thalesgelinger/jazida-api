package midw

import (
	"net/http"
	"os"
)

func WithLoaderAuth(handleFunc func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != os.Getenv("LOADER_PASS") {
			http.Error(w, "Only loaders can access", http.StatusForbidden)
			return
		}
		handleFunc(w, r)
	}
}
