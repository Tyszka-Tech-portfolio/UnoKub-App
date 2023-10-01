package web3

import (
	"net/http"
	"strings"
)

func EthAddressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := r.Header.Get("Eth-Address")
		if !strings.HasPrefix(address, "0x") || len(address) != 42 {
			http.Error(w, "Invalid Ethereum address", 400)
			return
		}
		next.ServeHTTP(w, r)
	})
}
