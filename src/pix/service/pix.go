package service

import (
	"mentoria/src/pix/model"
	"mentoria/src/pix/model/postgres"
)

type BankAccountRepo interface {
	CreateBankAccount(account *model.BankAccount) (*model.BankAccount, error)
	GetBankAccountById(id string) (*model.BankAccount, error)
}

type PixRepo interface {
	CreatePix(pix *model.Pix) (*model.Pix, error)
	GetPixById(id string) (*model.Pix, error)
	UpdatePixBalance(newPix *model.Pix) (*model.Pix, error)
}

type PixCodeRepo interface {
	CreatePixCode(pixCode *model.PixCode) (*model.PixCode, error)
	GetPixCodeByPixId(id string) (*model.PixCode, error)
	GetPixCodeByCode(code string) (*model.PixCode, error)
	DeletePixCode(code string) error
}

type TransactionRepo interface {
	CreateTransaction(newTransaction *model.Transaction) (*model.Transaction, error)
	ListUserTransactionsById(id string) ([]model.Transaction, error)
	UpdateTransactionById(transaction *model.Transaction) (*model.Transaction, error)
}
