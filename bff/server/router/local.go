package router

import (
	"net/http"
	"reflect"
)

type local struct {
	handler http.HandlerFunc
	method  string
	path    string
	model   reflect.Type
}

func (l *local) Handler() http.HandlerFunc {
	return l.handler
}

func (l *local) Method() string {
	return l.method
}

func (l *local) Path() string {
	return l.path
}

func (l *local) Model() reflect.Type {
	return l.model
}

func NewGetRoute(handler http.HandlerFunc, path string, model reflect.Type) Route {
	return NewRoute(handler, http.MethodGet, path, model)
}

func NewPostRoute(handler http.HandlerFunc, path string, model reflect.Type) Route {
	return NewRoute(handler, http.MethodPost, path, model)
}

// put ou patch? Revisar
func NewPutRoute(handler http.HandlerFunc, path string, model reflect.Type) Route {
	return NewRoute(handler, http.MethodPut, path, model)
}

func NewRoute(handler http.HandlerFunc, method string, path string, model reflect.Type) Route {
	return &local{
		handler: handler,
		method:  method,
		path:    path,
		model:   model,
	}
}
