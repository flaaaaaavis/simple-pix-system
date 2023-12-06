package user

import (
	gofrs "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"log"
	pb "mentoria/protobuf/user/v1"
	model "mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
	"mentoria/utils"
)

type Server struct {
	proto      pb.UnimplementedUserServiceServer
	userSvc    service.UserRepo
	contactSvc service.ContactRepo
}

func (s *Server) CreateUser(proto *pb.User) (*pb.User, error) {
	contactId := gofrs.UUID{}

	if proto.ContactId != "" {
		contactId = gofrs.FromStringOrNil(proto.ContactId)
	}

	newUser := &model.User{
		FullName:   proto.FullName,
		SocialName: proto.SocialName,
		CPF:        proto.Cpf,
		ContactID:  uuid.UUID(contactId),
	}

	res, err := s.userSvc.CreateUser(newUser)
	if err != nil {
		log.Fatalf("Error creating new user, %v", err)
		return nil, err
	}

	user := utils.FromUserModelToProto(res)

	return user, nil
}
