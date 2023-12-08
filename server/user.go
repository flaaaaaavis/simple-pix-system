package server

import (
	gofrs "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"log"
	pb "mentoria/protobuf/user/v1"
	model "mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
	"mentoria/utils"
)

type UserServer struct {
	proto      pb.UnimplementedUserServiceServer
	userSvc    service.UserRepo
	contactSvc service.ContactRepo
}

// CreateUser sends bff's user creation request to service through proto
// Criação do contato junto do usuário é obrigatória
func (s *UserServer) CreateUser(proto *pb.User) (*pb.User, error) {
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

// GetUserById sends bff's user obtaining request to service through proto
func (s *UserServer) GetUserById(id string) (*pb.User, error) {
	res, err := s.userSvc.GetUserById(id)
	if err != nil {
		log.Fatalf("Error getting user by id, %v", err)
		return nil, err
	}

	user := utils.FromUserModelToProto(res)

	return user, nil
}

// ListUsers sends bff's users listing request to service through proto
func (s *UserServer) ListUsers() ([]*pb.User, error) {
	modelUsers, err := s.userSvc.ListUsers()
	if err != nil {
		log.Fatalf("Error listing users, %v", err)
		return nil, err
	}

	var pbUsers []*pb.User

	for i := 0; i < len(modelUsers); i++ {
		pbUsers = append(pbUsers, utils.FromUserModelToProto(modelUsers[i]))
	}

	return pbUsers, nil
}

// UpdateUserById sends bff's user update request to service through proto
func (s *UserServer) UpdateUserById(proto *pb.User) (*pb.User, error) {
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

// Contact

// CreateContact sends bff's contact creation request to service through proto
func (s *UserServer) CreateContact(proto *pb.Contact) (*pb.Contact, error) {
	newContact := &model.Contact{
		Email:       proto.Email,
		PhoneNumber: proto.PhoneNumber,
	}

	res, err := s.contactSvc.CreateContact(newContact)
	if err != nil {
		log.Fatalf("Error creating new contact, %v", err)
		return nil, err
	}

	contact := utils.FromContactModelToProto(res)

	return contact, nil
}

// GetContactById sends bff's contact obtaining request to service through proto
func (s *UserServer) GetContactById(id string) (*pb.Contact, error) {
	res, err := s.contactSvc.GetContactById(id)
	if err != nil {
		log.Fatalf("Error getting contact by id, %v", err)
		return nil, err
	}

	user := utils.FromContactModelToProto(res)

	return user, nil
}

// UpdateContactById sends bff's contact update request to service through proto
func (s *UserServer) UpdateContactById(proto *pb.Contact) (*pb.Contact, error) {
	id := gofrs.FromStringOrNil(proto.Id)

	newContact := &model.Contact{
		ID:          uuid.UUID(id),
		Email:       proto.Email,
		PhoneNumber: proto.PhoneNumber,
	}

	res, err := s.contactSvc.UpdateContactById(newContact)
	if err != nil {
		log.Fatalf("Error updating contact, %v", err)
		return nil, err
	}

	contact := utils.FromContactModelToProto(res)

	return contact, nil
}
