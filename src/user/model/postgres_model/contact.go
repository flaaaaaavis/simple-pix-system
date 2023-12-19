package model

import (
	"github.com/google/uuid"
	pb "mentoria/protobuf/user/v1"
)

type Contact struct {
	ID          uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Email       string    `gorm:"type:varchar(255);column:email"`
	PhoneNumber string    `gorm:"type:varchar(255);column:phone_number"`
}

// GetContactByIdRequest model
type GetContactByIdRequest struct {
	ID string
}

// FromContactModelToProto converts contact from model to proto
func FromContactModelToProto(contact *Contact) *pb.Contact {
	return &pb.Contact{
		Id:          contact.ID.String(),
		Email:       contact.Email,
		PhoneNumber: contact.PhoneNumber,
	}
}

// GetContactByIdRequestFromProto converts GetContactByIdRequest from proto to model
func GetContactByIdRequestFromProto(req *pb.GetContactByIdRequest) *GetContactByIdRequest {
	return &GetContactByIdRequest{
		ID: req.Id,
	}
}
