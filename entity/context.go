package entity

import "net/http"

type Context struct {
	Body string
	Status int
	Request *http.Request
}

func NewContext(req *http.Request) *Context {
	return &Context{
		Request: req,
	}
}