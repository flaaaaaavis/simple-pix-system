package user

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
	"log"
	"mentoria/bff/types"
	pb "mentoria/protobuf/user/v1"
)

type grpc struct {
	client pb.UserServiceClient
}

func (g *grpc) CreateUser(ctx context.Context, request *types.CreateUserRequest) (*types.UserResponse, error) {
	newUser := &pb.User{
		FullName:   request.FullName,
		SocialName: request.SocialName,
		Cpf:        request.CPF,
		ContactId:  request.ContactID,
	}

	createdUser, err := g.client.CreateUser(ctx, newUser)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao criar novo usuário: %v", err)
		return nil, err
	}

	return &types.UserResponse{
		ID:         createdUser.Id,
		FullName:   createdUser.FullName,
		SocialName: createdUser.SocialName,
		CPF:        createdUser.Cpf,
		ContactID:  createdUser.ContactId,
	}, nil
}

func (g *grpc) GetUserById(ctx context.Context, request *types.GetUserByIdRequest) (*types.UserResponse, error) {
	newUser := &pb.GetUserByIdRequest{
		Id: request.ID,
	}

	createdUser, err := g.client.GetUserById(ctx, newUser)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao obter usuário pelo seu id: %v", err)
		return nil, err
	}

	return &types.UserResponse{
		ID:         createdUser.Id,
		FullName:   createdUser.FullName,
		SocialName: createdUser.SocialName,
		CPF:        createdUser.Cpf,
		ContactID:  createdUser.ContactId,
	}, nil
}

func (g *grpc) ListUsers(ctx context.Context) ([]*types.UserResponse, error) {
	res, err := g.client.ListUsers(ctx, &empty.Empty{})
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao listar usuários: %v", err)
		return nil, err
	}

	var newArr []*types.UserResponse

	for _, item := range res.Users {
		convertedItem := &types.UserResponse{
			ID:         item.Id,
			FullName:   item.FullName,
			SocialName: item.SocialName,
			CPF:        item.Cpf,
			ContactID:  item.ContactId,
		}
		newArr = append(newArr, convertedItem)
	}

	return newArr, nil
}

func (g *grpc) UpdateUserById(ctx context.Context, newUser *types.UpdateUserRequest) (*types.UserResponse, error) {
	reqUser := &pb.User{
		Id:         newUser.ID,
		FullName:   newUser.FullName,
		SocialName: newUser.SocialName,
		Cpf:        newUser.CPF,
		ContactId:  newUser.ContactID,
	}

	createdUser, err := g.client.UpdateUserById(ctx, reqUser)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao atualizar usuário: %v", err)
		return nil, err
	}

	return &types.UserResponse{
		ID:         createdUser.Id,
		FullName:   createdUser.FullName,
		SocialName: createdUser.SocialName,
		CPF:        createdUser.Cpf,
		ContactID:  createdUser.ContactId,
	}, nil
}

func (g *grpc) CreateContact(ctx context.Context, newContact *types.CreateContactRequest) (*types.ContactResponse, error) {
	newC := &pb.Contact{
		Email:       newContact.Email,
		PhoneNumber: newContact.PhoneNumber,
	}

	createdC, err := g.client.CreateContact(ctx, newC)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao criar novo contato: %v", err)
		return nil, err
	}

	return &types.ContactResponse{
		ID:          createdC.Id,
		Email:       createdC.Email,
		PhoneNumber: createdC.PhoneNumber,
	}, nil
}

func (g *grpc) GetContactById(ctx context.Context, req *types.GetContactByIdRequest) (*types.ContactResponse, error) {
	newC := &pb.GetContactByIdRequest{
		Id: req.ID,
	}

	createdC, err := g.client.GetContactById(ctx, newC)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao obter contato pelo seu id: %v", err)
		return nil, err
	}

	return &types.ContactResponse{
		ID:          createdC.Id,
		Email:       createdC.Email,
		PhoneNumber: createdC.PhoneNumber,
	}, nil
}

func (g *grpc) UpdateContactById(ctx context.Context, newContact *types.UpdateContactByIdRequest) (*types.ContactResponse, error) {
	newC := &pb.Contact{
		Id:          newContact.ID,
		Email:       newContact.Email,
		PhoneNumber: newContact.PhoneNumber,
	}

	createdC, err := g.client.UpdateContactById(ctx, newC)
	if err != nil {
		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}
		log.Fatalf("Erro ao atualizar contato: %v", err)
		return nil, err
	}

	return &types.ContactResponse{
		ID:          createdC.Id,
		Email:       createdC.Email,
		PhoneNumber: createdC.PhoneNumber,
	}, nil
}

func NewGRPC(user pb.UserServiceClient) Backend {
	return &grpc{
		client: user,
	}
}
