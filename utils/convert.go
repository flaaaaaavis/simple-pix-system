package utils

import (
	pixPb "mentoria/protobuf/pix/v1"
	userPb "mentoria/protobuf/user/v1"
	pixModel "mentoria/src/pix/model"
	model "mentoria/src/user/model/postgres_model"
)

// FromUserModelToProto converts user from model to proto
func FromUserModelToProto(user *model.User) *userPb.User {
	return &userPb.User{
		Id:         user.ID.String(),
		FullName:   user.FullName,
		SocialName: user.SocialName,
		Cpf:        user.CPF,
		ContactId:  user.ContactID.String(),
	}
}

// FromContactModelToProto converts contact from model to proto
func FromContactModelToProto(contact *model.Contact) *userPb.Contact {
	return &userPb.Contact{
		Id:          contact.ID.String(),
		Email:       contact.Email,
		PhoneNumber: contact.PhoneNumber,
	}
}

// FromBankAccountModelToProto converts BankAccount from model to proto
func FromBankAccountModelToProto(bankAccount *pixModel.BankAccount) *pixPb.BankAccount {
	return &pixPb.BankAccount{
		Id:            bankAccount.ID.String(),
		BankCode:      bankAccount.BankCode,
		BankName:      bankAccount.BankName,
		BankBranch:    bankAccount.BankBranch,
		AccountNumber: bankAccount.AccountNumber,
	}
}

// FromPixCodeModelToProto converts PixCode from model to proto
func FromPixCodeModelToProto(pixCode *pixModel.PixCode) *pixPb.PixCode {
	var pixType pixPb.PixType

	switch pixCode.Type {
	case pixModel.PixTypeEmail:
		pixType = pixPb.PixType_EMAIL
	case pixModel.PixTypePhone:
		pixType = pixPb.PixType_PHONE
	case pixModel.PixTypeCPF:
		pixType = pixPb.PixType_CPF
	case pixModel.PixTypeRandom:
		pixType = pixPb.PixType_RANDOM
	default:
		pixType = pixPb.PixType_RANDOM
	}

	return &pixPb.PixCode{
		Id:    pixCode.ID.String(),
		PixId: pixCode.PixID.String(),
		Type:  pixType,
		Code:  pixCode.Code,
	}
}

// FromPixModelToProto converts Pix from model to proto
func FromPixModelToProto(pix *pixModel.Pix) *pixPb.Pix {
	return &pixPb.Pix{
		Id:            pix.ID.String(),
		UserId:        pix.UserID.String(),
		BankAccountId: pix.BankAccountID.String(),
		Balance:       pix.Balance.String(),
	}
}

// FromTransactionModelToProto converts Transaction from model to proto
func FromTransactionModelToProto(transaction *pixModel.Transaction) *pixPb.Transaction {
	return &pixPb.Transaction{
		Id:         transaction.ID.String(),
		Date:       transaction.Date.String(),
		Amount:     transaction.Amount.String(),
		SenderId:   transaction.SenderID.String(),
		ReceiverId: transaction.ReceiverID.String(),
	}
}
