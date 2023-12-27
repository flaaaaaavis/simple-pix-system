package user

import (
	"context"
	"mentoria/bff/types"
)

type user interface {
	CreateUser(ctx context.Context, user *types.CreateUserRequest) (*types.UserResponse, error)
	GetUserById(ctx context.Context, req *types.GetUserByIdRequest) (*types.UserResponse, error)
	ListUsers(ctx context.Context) ([]*types.UserResponse, error)
	UpdateUserById(ctx context.Context, newUser *types.UpdateUserRequest) (*types.UserResponse, error)
}

type contact interface {
	CreateContact(ctx context.Context, newContact *types.CreateContactRequest) (*types.ContactResponse, error)
	GetContactById(ctx context.Context, req *types.GetContactByIdRequest) (*types.ContactResponse, error)
	UpdateContactById(ctx context.Context, newContact *types.UpdateContactByIdRequest) (*types.ContactResponse, error)
}

type Backend interface {
	user
	contact
}
