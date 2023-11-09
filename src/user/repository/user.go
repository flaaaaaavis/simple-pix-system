package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
	"projeto.com/src/user/service"
)

type userRepository struct {
	gormConnection *gorm.DB
}

func (u userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.gormConnection.Create(user)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new user: %v", err.Error)

		return nil, err.Error
	}

	return user, nil
}

func (u userRepository) GetUserById(id string) (*model.User, error) {
	user := &model.User{}
	condition := fmt.Sprintf("id=%v", id)

	result := u.gormConnection.First(model.User{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting user: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, user)
	if err != nil {
		fmt.Sprintf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return user, nil
}

func (u userRepository) ListUsers() ([]*model.User, error) {
	var users []*model.User

	result := u.gormConnection.Find(&users)
	if result.Error != nil {
		fmt.Sprintf("Error when listing users: %s", result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (u userRepository) UpdateUserById(newUser *model.User) (*model.User, error) {
	err := u.gormConnection.Model(newUser).Where("id IN (?)", newUser.ID).Updates(newUser)
	if err.Error != nil {
		fmt.Sprintf("Error when updating user: %v", err.Error)

		return nil, err.Error
	}

	return newUser, nil
}

func NewUser(db *gorm.DB) service.UserRepo {
	return &userRepository{
		gormConnection: db,
	}
}
