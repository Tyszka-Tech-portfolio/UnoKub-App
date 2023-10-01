package rate_limit

import (
	"net/http"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(25, 1)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}),
}
