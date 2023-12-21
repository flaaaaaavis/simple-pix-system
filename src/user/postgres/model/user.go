package model

import (
	pb "mentoria/protobuf/user/v1"
)

// User model
type User struct {
	ID         string `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	FullName   string `gorm:"type: varchar(255); column:full_name"`
	SocialName string `gorm:"type: varchar(255); column:social_name"`
	CPF        string `gorm:"type: varchar(255); column:cpf"`
	ContactID  string `gorm:"type:uuid;column:contact_id"`
}

// GetUserByIdRequest model
type GetUserByIdRequest struct {
	ID string
}

// FromUserModelToProto converts user from model to proto
func FromUserModelToProto(user *User) *pb.User {
	return &pb.User{
		Id:         user.ID,
		FullName:   user.FullName,
		SocialName: user.SocialName,
		Cpf:        user.CPF,
		ContactId:  user.ContactID,
	}
}

// GetUserByIdRequestFromProto converts GetUserByIdRequest from proto to model
func GetUserByIdRequestFromProto(req *pb.GetUserByIdRequest) *GetUserByIdRequest {
	return &GetUserByIdRequest{
		ID: req.Id,
	}
}
