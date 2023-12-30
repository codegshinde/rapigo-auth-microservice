package router

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func SetupRouter() *Router {
	return &Router{
		mux: http.DefaultServeMux,
	}
}

func (router *Router) GET(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.mux.HandleFunc(pattern, ApplyMiddleware(handler, middleware...))
}

func (router *Router) POST(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.mux.HandleFunc(pattern, ApplyMiddleware(handler, middleware...))
}
