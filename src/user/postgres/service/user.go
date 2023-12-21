package service

import (
	"context"
	model2 "mentoria/src/user/postgres/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model2.User) (*model2.User, error)
	GetUserById(ctx context.Context, req *model2.GetUserByIdRequest) (*model2.User, error)
	ListUsers(ctx context.Context) ([]*model2.User, error)
	UpdateUserById(ctx context.Context, newUser *model2.User) (*model2.User, error)
}

type ContactService interface {
	CreateContact(ctx context.Context, newContact *model2.Contact) (*model2.Contact, error)
	GetContactById(ctx context.Context, req *model2.GetContactByIdRequest) (*model2.Contact, error)
	UpdateContactById(ctx context.Context, newContact *model2.Contact) (*model2.Contact, error)
}
