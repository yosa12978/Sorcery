package sorcery

import (
	"net/http"
)

type Route struct {
	Pattern string
	Method  string
	Handler http.Handler
}

func NewRoute(pattern string, method string, handler http.Handler) Route {
	return Route{
		Pattern: pattern,
		Method:  method,
		Handler: handler,
	}
}

type Router struct {
	routes     []Route
	middleware []MiddlewareFunc
}

func NewRouter() *Router {
	return &Router{routes: []Route{}, middleware: []MiddlewareFunc{}}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range router.routes {
		if v.Pattern == r.URL.Path && v.Method == r.Method {
			v.Handler.ServeHTTP(w, r)
			return
		}
	}
	Http404handler.ServeHTTP(w, r)
}

type MiddlewareFunc func(next http.Handler) http.Handler

func (router *Router) Use(mw ...MiddlewareFunc) {
	router.middleware = append(router.middleware, mw...)
}

func (router *Router) GET(addr string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, "GET", handler))
}

func (router *Router) POST(addr string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, "POST", handler))
}

func (router *Router) PUT(addr string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, "PUT", handler))
}

func (router *Router) PATCH(addr string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, "PATCH", handler))
}

func (router *Router) DELETE(addr string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, "DELETE", handler))
}

func (router *Router) HandleSorc(addr string, method string, handler Sorc) {
	router.routes = append(router.routes, NewRoute(addr, method, handler))
}

func (router *Router) Handle(addr string, method string, handler http.Handler) {
	router.routes = append(router.routes, NewRoute(addr, method, handler))
}
