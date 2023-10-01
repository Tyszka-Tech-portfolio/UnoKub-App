package security

import (
	"net/http"
	"github.com/gorilla/csrf"
)

func CSRFProtection(next http.Handler) http.Handler {
	CSRF := csrf.Protect([]byte(**CSRF_SECRET_KEY**))
	return CSRF(next)
}
