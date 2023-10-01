package validation

import (
	"net/http"
)

func InputValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate input here
		// ...
		next.ServeHTTP(w, r)
	})
}
