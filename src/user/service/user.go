package service

import (
	"fmt"
	"projeto.com/src/user/model"
	"projeto.com/src/user/repository"
)

type UserInterface interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	ListUsers() ([]*model.User, error)
	UpdateUserById(newUser *model.User) (*model.User, error)
	CreateContact(contact *model.Contact) (*model.Contact, error)
	GetContactById(id string) (*model.Contact, error)
	UpdateContactById(newContact *model.Contact) (*model.Contact, error)
}

type userService struct {
	userRepo    repository.UserRepository
	contactRepo repository.ContactRepository
}

func (u *userService) CreateUser(user model.User) (*model.User, error) {
	newUser, err := u.CreateUser(user)
	if err != nil {
		fmt.Sprintf("Error when creating new user on user service: %s", err)

		return nil, err
	}
	return newUser, nil
}

func (u *userService) GetUserById(id string) (*model.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		fmt.Sprintf("Error when getting user from id on user service: %s", err)

		return nil, err
	}
	return user, nil
}

func (u *userService) ListUsers() ([]*model.User, error) {
	users, err := u.ListUsers()
	if err != nil {
		fmt.Sprintf("Error when listing users on user service: %s", err)

		return nil, err
	}
	return users, nil
}

func (u *userService) UpdateUserById(newUser *model.User) (*model.User, error) {
	user, err := u.UpdateUserById(newUser)
	if err != nil {
		fmt.Sprintf("Error when updating user on user service: %s", err)

		return nil, err
	}
	return user, nil
}

func (u *userService) CreateContact(contact *model.Contact) (*model.Contact, error) {
	newContact, err := u.CreateContact(contact)
	if err != nil {
		fmt.Sprintf("Error when creating new contact on user service: %s", err)

		return nil, err
	}
	return newContact, nil
}

func (u *userService) GetContactById(id string) (*model.Contact, error) {
	contact, err := u.GetContactById(id)
	if err != nil {
		fmt.Sprintf("Error when getting contact from id on user service: %s", err)

		return nil, err
	}
	return contact, nil
}

func (u *userService) UpdateContactById(newContact *model.Contact) (*model.Contact, error) {
	contact, err := u.UpdateContactById(newContact)
	if err != nil {
		fmt.Sprintf("Error when getting contact from id on user service: %s", err)

		return nil, err
	}
	return contact, nil
}

func NewUserService(userRepo repository.UserRepository, contactRepo repository.ContactRepository) UserInterface {
	return &userService{
		userRepo:    userRepo,
		contactRepo: contactRepo,
	}
}
