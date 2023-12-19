package model

import (
	"github.com/google/uuid"
	pb "mentoria/protobuf/pix/v1"
)

// BankAccount model
type BankAccount struct {
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	BankCode      string    `json:"bank_code"`
	BankName      string    `json:"bank_name"`
	BankBranch    string    `json:"bank_branch"`
	AccountNumber string    `json:"account_number"`
}

// FromBankAccountModelToProto converts BankAccount from model to proto
func FromBankAccountModelToProto(bankAccount *BankAccount) *pb.BankAccount {
	return &pb.BankAccount{
		Id:            bankAccount.ID.String(),
		BankCode:      bankAccount.BankCode,
		BankName:      bankAccount.BankName,
		BankBranch:    bankAccount.BankBranch,
		AccountNumber: bankAccount.AccountNumber,
	}
}
