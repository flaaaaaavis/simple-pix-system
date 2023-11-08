package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"projeto.com/src/config"
	"projeto.com/src/user/model"
)

type UserRepository struct {
	gormConnection config.Gorm
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	db := u.gormConnection.Db()

	id := uuid.New().String()

	newUser := &model.User{
		ID:         id,
		FullName:   user.FullName,
		SocialName: user.SocialName,
		CPF:        user.CPF,
		ContactID:  user.ContactID,
	}

	err := db.Create(newUser)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new user: %v", err.Error)

		return nil, err.Error
	}

	return newUser, nil
}

func (u *UserRepository) GetUserById(id string) (*model.User, error) {
	db := u.gormConnection.Db()

	user := &model.User{}

	condition := fmt.Sprintf("id=%v", id)

	result := db.First(model.User{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting user: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err.Error != nil {
		fmt.Sprintf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, user)
	if err.Error != nil {
		fmt.Sprintf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return user, nil
}

func (u *UserRepository) ListUsers() ([]*model.User, error) {
	return nil, nil
}

func (u *UserRepository) UpdateUserById(newUser *model.User) (*model.User, error) {
	db := u.gormConnection.Db()

	user := &model.User{
		FullName:   newUser.FullName,
		SocialName: newUser.SocialName,
		CPF:        newUser.CPF,
		ContactID:  newUser.ContactID,
	}

	err := db.Model(user).Where("id IN (?)", newUser.ID).Updates(user)
	if err.Error != nil {
		fmt.Sprintf("Error when updating user: %v", err.Error)

		return nil, err.Error
	}

	return user, nil
}
