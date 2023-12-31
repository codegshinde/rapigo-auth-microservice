package routes

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func SetupRouter() *Router {
	return &Router{
		mux: http.DefaultServeMux,
	}
}

func (router *Router) GET(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc(pattern, handler)
}

func (router *Router) POST(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc(pattern, handler)
}
