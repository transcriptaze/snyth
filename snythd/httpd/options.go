package httpd

import (
	"net/http"
)

func (d dispatcher) options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

	if d.debug {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	} else {
		w.Header().Set("Cache-Control", "max-age=3600, must-revalidate")
	}
}
