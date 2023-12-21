package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"mentoria/src/user/postgres/model"
	"mentoria/src/user/postgres/service"
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

	id := fmt.Sprintf("%v", req.ID)
	condition := &model.User{
		ID: id,
	}

	u.gormConnection.First(user, condition)

	return user, nil
}

func (u userRepository) ListUsers(_ context.Context) ([]*model.User, error) {
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
