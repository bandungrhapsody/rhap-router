package fw

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	prefixPath  string
	routes      []RouteEntry
	middlewares []Middleware
}

type Middleware func(handler Handler) Handler

func NewRouter() *Router {
	return &Router{}
}

func (rtr *Router) Prefix(prefix string) *Router {
	rtr.prefixPath = prefix
	return rtr
}

func (rtr *Router) GET(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodGet, path, handler)
}

func (rtr *Router) POST(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodPost, path, handler)
}

func (rtr *Router) PUT(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodPut, path, handler)
}

func (rtr *Router) DELETE(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodDelete, path, handler)
}

func (rtr *Router) Routes(path string, fn MethodSetter) {
	gr := &GroupRoutes{path: rtr.prefixPath + path}

	fn(gr)

	if gr.routes != nil {
		rtr.routes = append(rtr.routes, gr.routes...)
	}
}

func (rtr *Router) Use(m ...Middleware) {
	rtr.middlewares = append(rtr.middlewares, m...)
}

func (rtr *Router) Listen(port int) {
	strPort := fmt.Sprintf(":%d", port)

	log.Fatal(http.ListenAndServe(strPort, rtr))
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rtr.routes {
		params := route.match(r)
		if params == nil {
			continue
		}

		rtr.serve(route, &Context{
			writer:  w,
			request: r,
			params:  params,
		})
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) serve(route RouteEntry, ctx *Context) {
	if len(rtr.middlewares) < 1 {
		route.HandlerFunc(ctx)
		return
	}

	wrapped := route.HandlerFunc
	for i := len(rtr.middlewares) - 1; i >= 0; i-- {
		wrapped = rtr.middlewares[i](wrapped)
	}
	wrapped(ctx)
}

func (rtr *Router) addRouteEntry(method, path string, handler Handler) {
	path = rtr.prefixPath + path
	exactPath := generatePath(path)

	rtr.routes = append(rtr.routes, RouteEntry{
		Method: method,
		Path: exactPath,
		HandlerFunc: handler,
	})
}
