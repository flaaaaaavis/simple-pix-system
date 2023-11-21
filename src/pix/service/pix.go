package service

import (
	"mentoria/src/pix/model/postgres"
)

type BankAccountRepo interface {
	CreateBankAccount(account *postgres.BankAccount) (*postgres.BankAccount, error)
	GetBankAccountById(id string) (*postgres.BankAccount, error)
}

type PixRepo interface {
	CreatePix(pix *postgres.Pix) (*postgres.Pix, error)
	GetPixById(id string) (*postgres.Pix, error)
	UpdatePixBalance(newPix *postgres.Pix) (*postgres.Pix, error)
}

type PixCodeRepo interface {
	CreatePixCode(pixCode *postgres.PixCode) (*postgres.PixCode, error)
	GetPixCodeByPixId(id string) (*postgres.PixCode, error)
	GetPixCodeByCode(code string) (*postgres.PixCode, error)
	DeletePixCode(code string) error
}

type TransactionRepo interface {
	CreateTransaction(newTransaction *postgres.Transaction) (*postgres.Transaction, error)
	ListUserTransactionsById(id string) ([]postgres.Transaction, error)
	UpdateTransactionById(transaction *postgres.Transaction) (*postgres.Transaction, error)
}
