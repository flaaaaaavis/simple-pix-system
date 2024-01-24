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
	routes = append(routes, router.NewPostRoute(r.CreateUser, "/create", reflect.TypeOf(types.CreateUserRequest{})))

	// user routes
	routes = append(routes, router.NewPostRoute(r.CreateUser, "/create-user", reflect.TypeOf(types.CreateUserRequest{})))
	routes = append(routes, router.NewGetRoute(r.GetUserById, "/get-user", reflect.TypeOf(types.GetUserByIdRequest{})))
	routes = append(routes, router.NewGetRoute(r.ListUsers, "/list-users", reflect.TypeOf(nil)))
	routes = append(routes, router.NewPostRoute(r.UpdateUserById, "/update-user", reflect.TypeOf(types.UpdateUserRequest{})))

	// contact routes
	routes = append(routes, router.NewPostRoute(r.CreateContact, "/create-contact", reflect.TypeOf(types.CreateContactRequest{})))
	routes = append(routes, router.NewPostRoute(r.GetContactById, "/get-contact", reflect.TypeOf(types.GetContactByIdRequest{})))
	routes = append(routes, router.NewPostRoute(r.UpdateContactById, "/update-contact", reflect.TypeOf(types.UpdateContactByIdRequest{})))

	return routes
}

func NewRouter(backend Backend) router.Router {
	return &RouterImplementation{
		backend: backend,
	}
}
