package model

import (
	"github.com/shopspring/decimal"
	pb "mentoria/protobuf/pix/v1"
)

// Pix model
type Pix struct {
	ID            string          `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	UserID        string          `gorm:"type:uuid;column:user_id"`
	BankAccountID string          `gorm:"type:uuid;column:bank_account_id"`
	Balance       decimal.Decimal `gorm:"type:uuid;column:balance"`
}

// FromPixModelToProto converts Pix from model to proto
func FromPixModelToProto(pix *Pix) *pb.Pix {
	return &pb.Pix{
		Id:            pix.ID,
		UserId:        pix.UserID,
		BankAccountId: pix.BankAccountID,
		Balance:       pix.Balance.String(),
	}
}
