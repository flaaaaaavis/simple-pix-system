package model

import (
	pb "mentoria/protobuf/pix/v1"
)

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

// PixCode model
type PixCode struct {
	ID    string  `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	PixID string  `gorm:"type:uuid;column:pix_id"`
	Type  PixType `gorm:"type: varchar(255);column:type"`
	Code  string  `gorm:"type: varchar(255);column:code"`
}

// FromPixCodeModelToProto converts PixCode from model to proto
func FromPixCodeModelToProto(pixCode *PixCode) *pb.PixCode {
	var pixType pb.PixType

	switch pixCode.Type {
	case PixTypeEmail:
		pixType = pb.PixType_EMAIL
	case PixTypePhone:
		pixType = pb.PixType_PHONE
	case PixTypeCPF:
		pixType = pb.PixType_CPF
	case PixTypeRandom:
		pixType = pb.PixType_RANDOM
	default:
		pixType = pb.PixType_RANDOM
	}

	return &pb.PixCode{
		Id:    pixCode.ID,
		PixId: pixCode.PixID,
		Type:  pixType,
		Code:  pixCode.Code,
	}
}
