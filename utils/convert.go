package utils

import (
	pb "mentoria/protobuf/user/v1"
	model "mentoria/src/user/model/postgres_model"
)

func FromUserModelToProto(user *model.User) *pb.User {
	contactId := ""
	if user.ContactID.String() != "" {
		contactId = user.ContactID.String()
	}

	return &pb.User{
		Id:         user.ID.String(),
		FullName:   user.FullName,
		SocialName: user.SocialName,
		Cpf:        user.CPF,
		ContactId:  contactId,
	}
}
