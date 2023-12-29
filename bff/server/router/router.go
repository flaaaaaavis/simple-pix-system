package router

import (
	"net/http"
	"reflect"
)

type Route interface {
	Handler() http.HandlerFunc
	Method() string
	Path() string
	Model() reflect.Type
}

type Router interface {
	Routes() []Route
}
