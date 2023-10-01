package cors

import (
	"net/http"
)

func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", **CORS_ALLOWED_ORIGINS**)
		next.ServeHTTP(w, r)
	})
}
