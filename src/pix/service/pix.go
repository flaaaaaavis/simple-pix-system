package service

import (
	"mentoria/src/pix/model"
)

type BankAccountService interface {
	CreateBankAccount(account *model.BankAccount) (*model.BankAccount, error)
	GetBankAccountById(id string) (*model.BankAccount, error)
}

type PixService interface {
	CreatePix(pix *model.Pix) (*model.Pix, error)
	GetPixById(id string) (*model.Pix, error)
	UpdatePixBalance(newPix *model.Pix) (*model.Pix, error)
}

type PixCodeService interface {
	CreatePixCode(pixCode *model.PixCode) (*model.PixCode, error)
	GetPixCodeByPixId(id string) (*model.PixCode, error)
	GetPixCodeByCode(code string) (*model.PixCode, error)
	DeletePixCode(code string) error
}

type TransactionService interface {
	CreateTransaction(newTransaction *model.Transaction) (*model.Transaction, error)
	ListUserTransactionsById(id string) ([]model.Transaction, error)
}
