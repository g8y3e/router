package router

import (
	"github.com/g8y3e/router/controller"
	"github.com/g8y3e/router/entity"
	"net/http"
	"regexp"
)

type Router struct {
	prefix string

	getRoutes map[string]*Route
	postRoutes map[string]*Route
	putRoutes map[string]*Route
	deleteRoutes map[string]*Route

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
		getRoutes: map[string]*Route{},
		postRoutes: map[string]*Route{},
		putRoutes: map[string]*Route{},
		deleteRoutes: map[string]*Route{},
	}
}

func (r *Router) Get(path string) *Route {
	pathRoute := NewRoute()
	r.getRoutes[path] = pathRoute
	return pathRoute
}

func (r *Router) Post(path string) *Route {
	pathRoute := NewRoute()
	r.postRoutes[path] = pathRoute
	return pathRoute
}

func (r *Router) Put(path string) *Route {
	pathRoute := NewRoute()
	r.putRoutes[path] = pathRoute
	return pathRoute
}

func (r *Router) Delete(path string) *Route {
	pathRoute := NewRoute()
	r.deleteRoutes[path] = pathRoute
	return pathRoute
}

func (r *Router) Match(req *http.Request) *Route {
	// check routes type
	var searchRoutes map[string]*Route
	if req.Method == http.MethodGet {
		searchRoutes = r.getRoutes
	} else if req.Method == http.MethodPost {
		searchRoutes = r.postRoutes
	} else if req.Method == http.MethodPut {
		searchRoutes = r.putRoutes
	} else if req.Method == http.MethodDelete {
		searchRoutes = r.deleteRoutes
	} else if searchRoutes == nil {
		return nil
	}

	for key, value := range searchRoutes {
		reg := regexp.MustCompile(r.prefix + key)
		if reg.MatchString(req.URL.Path) {
			return value
		}
	}
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
