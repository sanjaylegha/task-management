// backend/middleware.go

package api

import (
	"net/http"

	"github.com/rs/cors"
)

// CorsMiddleware adds CORS headers to the response
func CorsMiddleware(next http.Handler) http.Handler {
	c := cors.Default()
	return c.Handler(next)
}
