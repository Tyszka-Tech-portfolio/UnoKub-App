package main

import (
	"log"
	"net/http"

	"github.com/Tyszka-Tech-portfolio/UnoKub-App/api" // Replace with the actual path to your api package
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/auth"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/cache"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/cors"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/logging"
	monitoring "github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/monitioring"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/rate_limit"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/security"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/validation"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/middleware/web3"
	"github.com/Tyszka-Tech-portfolio/UnoKub-App/utils" // Replace with the actual path to your utils package
	"github.com/justinas/alice"
)

func main() {
	// Initialize the database
	utils.InitDB()

	// Initialize your routes
	r := api.NewRouter()

	// Initialize Alice middleware chaining
	chain := alice.New(
		logging.RequestLogger,
		security.CSRFProtection,
		monitoring.MetricsMiddleware,
		rate_limit.RateLimiter,
		cors.CORSHandler,
		auth.JWTValidator,
		auth.SessionManager,
		web3.JSONRPCMiddleware,
		web3.EthAddressMiddleware,
		web3.SignedMessageAuthMiddleware,
		cache.CacheHandler,
		validation.InputValidator,
	).Then(r)
	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", chain))
}
