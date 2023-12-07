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

// CreateUser sends bff's user creation request to service thru proto
// Criação do contato junto do usuário é obrigatória
func (s *Server) CreateUser(proto *pb.User) (*pb.User, error) {
	contactId := gofrs.FromStringOrNil(proto.ContactId)

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

// GetUserById sends bff's user obtaining request to service thru proto
func (s *Server) GetUserById(id string) (*pb.User, error) {
	res, err := s.userSvc.GetUserById(id)
	if err != nil {
		log.Fatalf("Error getting user by id, %v", err)
		return nil, err
	}

	user := utils.FromUserModelToProto(res)

	return user, nil
}

// ListUsers sends bff's users listing request to service thru proto
func (s *Server) ListUsers() ([]*pb.User, error) {
	modelUsers, err := s.userSvc.ListUsers()
	if err != nil {
		log.Fatalf("Error listing users, %v", err)
		return nil, err
	}

	pbUsers := []*pb.User{}

	for i := 0; i < len(modelUsers); i++ {
		pbUsers = append(pbUsers, utils.FromUserModelToProto(modelUsers[i]))
	}

	return pbUsers, nil
}

// UpdateUserById sends bff's user update request to service thru proto
func (s *Server) UpdateUserById(proto *pb.User) (*pb.User, error) {
	id := gofrs.FromStringOrNil(proto.Id)
	contactId := gofrs.FromStringOrNil(proto.ContactId)

	newUser := &model.User{
		ID:         uuid.UUID(id),
		FullName:   proto.FullName,
		SocialName: proto.SocialName,
		CPF:        proto.Cpf,
		ContactID:  uuid.UUID(contactId),
	}

	res, err := s.userSvc.UpdateUserById(newUser)
	if err != nil {
		log.Fatalf("Error updating user, %v", err)
		return nil, err
	}

	user := utils.FromUserModelToProto(res)

	return user, nil
}
