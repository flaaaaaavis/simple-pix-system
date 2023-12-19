package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	pb "mentoria/protobuf/pix/v1"
)

// Pix model
type Pix struct {
	ID            uuid.UUID       `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	UserID        uuid.UUID       `gorm:"type:uuid;column:user_id"`
	BankAccountID uuid.UUID       `gorm:"type:uuid;column:bank_account_id"`
	Balance       decimal.Decimal `gorm:"type:uuid;column:balance"`
}

// FromPixModelToProto converts Pix from model to proto
func FromPixModelToProto(pix *Pix) *pb.Pix {
	return &pb.Pix{
		Id:            pix.ID.String(),
		UserId:        pix.UserID.String(),
		BankAccountId: pix.BankAccountID.String(),
		Balance:       pix.Balance.String(),
	}
}
