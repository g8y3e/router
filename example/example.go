package main

import (
	"fmt"
	"github.com/g8y3e/router"
	"net/http"
)

type TestController struct {
}

func (t * TestController) Process(w http.ResponseWriter, req *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"version": "%s"}`, "1.0.0")
	return nil
}

func main() {
	cf := &router.Config{
		Prefix:"/test",
	}

	r := router.New(cf)

	// 1 variant
	r.Get("/test1").Controller(&TestController{})

	// 2 variant
	testRoute := r.Get("/test2")
	testRoute.Middleware(&TestController{}, &TestController{})
	testRoute.Controller(&TestController{})

	port := 8080

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
