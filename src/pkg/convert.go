package pkg

import (
	pb "mentoria/protos/protobuf/user/v1"
	model "mentoria/src/user/model/postgres_model"
)

func ContatcFromProto(c *model.Contact) *pb.Contact {
	return &pb.Contact{
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
}
