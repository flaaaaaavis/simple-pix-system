package server

import (
	"context"
	gofrs "github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"log"
	pb "mentoria/protobuf/user/v1"
	model "mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

// UserServer communication interface between bff and proto
type UserServer struct {
	pb.UnimplementedUserServiceServer
	UserSvc    service.UserService
	ContactSvc service.ContactService
}

// CreateUser sends bff's user creation request to service through proto
// Criação do contato junto do usuário é obrigatória
func (s *UserServer) CreateUser(ctx context.Context, proto *pb.User) (*pb.User, error) {
	contactId := gofrs.FromStringOrNil(proto.ContactId)

	newUser := &model.User{
		FullName:   proto.FullName,
		SocialName: proto.SocialName,
		CPF:        proto.Cpf,
		ContactID:  uuid.UUID(contactId),
	}

	res, err := s.UserSvc.CreateUser(ctx, newUser)
	if err != nil {
		log.Fatalf("Error creating new user, %v", err)
		return nil, err
	}

	user := model.FromUserModelToProto(res)

	return user, nil
}

// GetUserById sends bff's user obtaining request to service through proto
func (s *UserServer) GetUserById(ctx context.Context, proto *pb.GetUserByIdRequest) (*pb.User, error) {
	modelReq := model.GetUserByIdRequestFromProto(proto)
	res, err := s.UserSvc.GetUserById(ctx, modelReq)
	if err != nil {
		log.Fatalf("Error getting user by id, %v", err)
		return nil, err
	}

	user := model.FromUserModelToProto(res)

	return user, nil
}

// ListUsers sends bff's users listing request to service through proto
func (s *UserServer) ListUsers(ctx context.Context, _ *empty.Empty) (*pb.ListUsersResponse, error) {
	modelUsers, err := s.UserSvc.ListUsers(ctx)
	if err != nil {
		log.Fatalf("Error listing users, %v", err)
		return nil, err
	}

	var pbUsers []*pb.User

	for i := 0; i < len(modelUsers); i++ {
		pbUsers = append(pbUsers, model.FromUserModelToProto(modelUsers[i]))
	}

	listRes := &pb.ListUsersResponse{
		Users: pbUsers,
	}

	return listRes, nil
}

// UpdateUserById sends bff's user update request to service through proto
func (s *UserServer) UpdateUserById(ctx context.Context, proto *pb.User) (*pb.User, error) {
	id := gofrs.FromStringOrNil(proto.Id)
	contactId := gofrs.FromStringOrNil(proto.ContactId)

	newUser := &model.User{
		ID:         uuid.UUID(id),
		FullName:   proto.FullName,
		SocialName: proto.SocialName,
		CPF:        proto.Cpf,
		ContactID:  uuid.UUID(contactId),
	}

	res, err := s.UserSvc.UpdateUserById(ctx, newUser)
	if err != nil {
		log.Fatalf("Error updating user, %v", err)
		return nil, err
	}

	user := model.FromUserModelToProto(res)

	return user, nil
}

// Contact

// CreateContact sends bff's contact creation request to service through proto
func (s *UserServer) CreateContact(ctx context.Context, proto *pb.Contact) (*pb.Contact, error) {
	newContact := &model.Contact{
		Email:       proto.Email,
		PhoneNumber: proto.PhoneNumber,
	}

	res, err := s.ContactSvc.CreateContact(ctx, newContact)
	if err != nil {
		log.Fatalf("Error creating new contact, %v", err)
		return nil, err
	}

	contact := model.FromContactModelToProto(res)

	return contact, nil
}

// GetContactById sends bff's contact obtaining request to service through proto
func (s *UserServer) GetContactById(ctx context.Context, proto *pb.GetContactByIdRequest) (*pb.Contact, error) {
	modelReq := model.GetContactByIdRequestFromProto(proto)
	res, err := s.ContactSvc.GetContactById(ctx, modelReq)
	if err != nil {
		log.Fatalf("Error getting contact by id, %v", err)
		return nil, err
	}

	user := model.FromContactModelToProto(res)

	return user, nil
}

// UpdateContactById sends bff's contact update request to service through proto
func (s *UserServer) UpdateContactById(ctx context.Context, proto *pb.Contact) (*pb.Contact, error) {
	id := gofrs.FromStringOrNil(proto.Id)

	newContact := &model.Contact{
		ID:          uuid.UUID(id),
		Email:       proto.Email,
		PhoneNumber: proto.PhoneNumber,
	}

	res, err := s.ContactSvc.UpdateContactById(ctx, newContact)
	if err != nil {
		log.Fatalf("Error updating contact, %v", err)
		return nil, err
	}

	contact := model.FromContactModelToProto(res)

	return contact, nil
}
