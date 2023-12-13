package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

type userRepository struct {
	gormConnection *gorm.DB
}

func (u userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.gormConnection.Create(user)
	if err.Error != nil {
		log.Fatalf("Error when creating new user: %v", err.Error)

		return nil, err.Error
	}

	return user, nil
}

func (u userRepository) GetUserById(id string) (*model.User, error) {
	user := &model.User{}
	condition := fmt.Sprintf("id=%v", id)

	result := u.gormConnection.First(model.User{}, condition)
	if result.Error != nil {
		log.Fatalf("Error when getting user: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		log.Fatalf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, user)
	if err != nil {
		log.Fatalf("Error when getting user: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return user, nil
}

func (u userRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	result := u.gormConnection.Find(&users)
	if result.Error != nil {
		log.Fatalf("Error when listing users: %s", result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (u userRepository) UpdateUserById(newUser *model.User) (*model.User, error) {
	err := u.gormConnection.Model(newUser).Where("id IN (?)", newUser.ID).Updates(newUser)
	if err.Error != nil {
		log.Fatalf("Error when updating user: %v", err.Error)

		return nil, err.Error
	}

	return newUser, nil
}

func NewUser(db *gorm.DB) service.UserRepo {
	return &userRepository{
		gormConnection: db,
	}
}
