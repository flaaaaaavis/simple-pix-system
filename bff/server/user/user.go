package user

import (
	"mentoria/bff/server/router"
	"mentoria/bff/types"
	"reflect"
)

type RouterImplementation struct {
	backend Backend
}

func (r *RouterImplementation) Routes() []router.Route {
	routes := make([]router.Route, 0)
	routes = append(routes, router.NewPostRoute(r.CreateUser, "uau", reflect.TypeOf(types.CreateUserRequest{})))

	return routes
}

func NewRouter(backend Backend) router.Router {
	return &RouterImplementation{
		backend: backend,
	}
}
