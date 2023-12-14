package user

import "mentoria/api/types"

type user interface {
	CreateUserRequest(types types.User) types.User
}

type Backend interface {
	user
}
