package monitoring

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == **METRICS_ENDPOINT** {
			promhttp.Handler().ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
