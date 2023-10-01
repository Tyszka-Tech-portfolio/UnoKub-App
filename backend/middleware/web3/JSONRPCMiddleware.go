package web3

import (
	"encoding/json"
	"net/http"
)

func JSONRPCMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON-RPC request", 400)
			return
		}
		if _, ok := payload["jsonrpc"]; !ok {
			http.Error(w, "Missing jsonrpc field", 400)
			return
		}
		next.ServeHTTP(w, r)
	})
}
