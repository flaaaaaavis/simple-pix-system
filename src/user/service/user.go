package service

import (
	"projeto.com/src/user/model"
)

type UserRepo interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	ListUsers() ([]*model.User, error)
	UpdateUserById(newUser *model.User) (*model.User, error)
}

type ContactRepo interface {
	CreateContact(contact *model.Contact) (*model.Contact, error)
	GetContactById(id string) (*model.Contact, error)
	UpdateContactById(newContact *model.Contact) (*model.Contact, error)
}