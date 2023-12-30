package router

import "net/http"

// Middleware type for middleware functions
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddleware applies a slice of middleware to a given handler
func ApplyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
