package router

import (
	"github.com/g8y3e/router/entity"
	"net/http"
)

type Route struct {
	path string
	middleware []entity.IController
	pathController entity.IController
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) Middleware(controllers ...entity.IController) *Route {
	for _, c := range controllers {
		r.middleware = append(r.middleware, c)
	}
	return r
}

func (r *Route) Controller(controller entity.IController) *Route {
	r.pathController = controller
	return r
}

func (r *Route) Process(w http.ResponseWriter, req *http.Request) {
	// process middleware
	for _, m := range r.middleware {
		err := m.Process(w, req)
		if err != nil {
			return
		}
	}

	// process controllers
	r.pathController.Process(w, req)
}
