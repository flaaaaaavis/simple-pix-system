package repository

import (
	"projeto.com/src/config"
	"projeto.com/src/pix/model"
)

type TransactionRepository struct {
	gormConnection config.Gorm
}

func CreateTransaction() model.Transaction {
	return model.Transaction{}
}

func ListUserTransactionsById(id string) []model.Transaction {

	return []model.Transaction{}
}

func updateTransaction() model.Transaction {
	return model.Transaction{}
}
