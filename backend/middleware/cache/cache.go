package cache

import (
	"net/http"
	"time"
)

func CacheHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age="+**CACHE_DURATION**)
		next.ServeHTTP(w, r)
	})
}
