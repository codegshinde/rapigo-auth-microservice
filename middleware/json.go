package middleware

import (
	"net/http"
)

// JSONMiddleware is a middleware that sets the Content-Type header to application/json
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
