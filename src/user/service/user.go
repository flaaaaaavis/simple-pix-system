package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"mentoria/src/user/model/postgres_model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserById(ctx context.Context, req *model.GetUserByIdRequest) (*model.User, error)
	ListUsers(ctx context.Context, empty *empty.Empty) ([]*model.User, error)
	UpdateUserById(ctx context.Context, newUser *model.User) (*model.User, error)
}

type ContactRepo interface {
	CreateContact(ctx context.Context, newContact *model.Contact) (*model.Contact, error)
	GetContactById(ctx context.Context, req *model.GetContactByIdRequest) (*model.Contact, error)
	UpdateContactById(ctx context.Context, newContact *model.Contact) (*model.Contact, error)
}
