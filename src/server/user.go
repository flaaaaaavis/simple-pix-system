package server

import (
	"context"
	"log"

	pb "mentoria/protos/protobuf/user/v1"
	"mentoria/src/pkg"
	model "mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	svcRepo service.ContactRepo
}

func (u *UserService) CreateContact(ctx context.Context, c *pb.Contact) (*pb.Contact, error) {

	modelUser := &model.Contact{
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
	res, err := u.svcRepo.CreateContact(modelUser)

	if err != nil {
		log.Fatalf("error on create contact, %s", err)
		return nil, err
	}

	conv := pkg.ContatcFromProto(res)

	return conv, nil
}
