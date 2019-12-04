package router

import (
	"github.com/g8y3e/router/controller"
	"github.com/g8y3e/router/entity"
	"github.com/g8y3e/router/route"
	"net/http"
)

type Router struct {
	prefix string
	routes map[string]*route.Route

	httpNotFound entity.IController
}

func New(cf *Config) *Router {
	httpNotFound := cf.HttpNotFound
	if httpNotFound == nil {
		httpNotFound = &controller.HttpNotFound{}
	}

	return &Router{
		prefix: cf.Prefix,
		httpNotFound: httpNotFound,
	}
}

func (r *Router) Get(path string) *route.Route {
	pathRoute := route.New()
	r.routes[path] = pathRoute
	return pathRoute
}

func (r *Router) Match(req *http.Request) *route.Route {
	return nil
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	routeMatch := r.Match(req)
	if routeMatch == nil {
		r.httpNotFound.Process(w, req)
	} else {
		routeMatch.Process(w, req)
	}
}
