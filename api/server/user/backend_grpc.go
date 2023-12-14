package user

import (
	"context"

	"mentoria/api/types"
	pb "mentoria/protos/protobuf/user/v1"
)

type grpc struct {
	user pb.UserServiceClient
}

func (g *grpc) CreateUser(ctx context.Context, body types.User) (*types.User, error)  {

	req := &pb.User{
		FullName:   body.FullName,
		SocialName: body.SocialName,
		Cpf:        body.Cpf,
		ContactId:  body.ContactID,
	}

	res, err := g.user.CreateUser(ctx, req)

	if err != nil{
		return nil,err
	}

	
	return res, nil


}
