package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
	"log"
	"mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

type userRepository struct {
	gormConnection *gorm.DB
}

func (u userRepository) CreateUser(_ context.Context, user *model.User) (*model.User, error) {
	err := u.gormConnection.Create(user)
	if err.Error != nil {
		log.Fatalf("Error when creating new user: %v", err.Error)

		return nil, err.Error
	}

	return user, nil
}

func (u userRepository) GetUserById(_ context.Context, req *model.GetUserByIdRequest) (*model.User, error) {
	user := &model.User{}

	id := fmt.Sprintf("\"%v\"", req.ID)
	condition := fmt.Sprintf("\"id\"=%v", id)

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

func (u userRepository) ListUsers(_ context.Context, _ *empty.Empty) ([]*model.User, error) {
	var users []*model.User

	result := u.gormConnection.Find(&users)
	if result.Error != nil {
		log.Fatalf("Error when listing users: %s", result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (u userRepository) UpdateUserById(_ context.Context, newUser *model.User) (*model.User, error) {
	err := u.gormConnection.Model(newUser).Where("id IN (?)", newUser.ID).Updates(newUser)
	if err.Error != nil {
		log.Fatalf("Error when updating user: %v", err.Error)

		return nil, err.Error
	}

	return newUser, nil
}

func NewUser(db *gorm.DB) service.UserService {
	return &userRepository{
		gormConnection: db,
	}
}
