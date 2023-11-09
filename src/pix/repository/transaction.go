package repository

import (
	"gorm.io/gorm"
	"projeto.com/src/pix/model"
)

type TransactionRepository struct {
	gormConnection *gorm.DB
}

func CreateTransaction(newTransaction *model.Transaction) (*model.Transaction, error) {
	return &model.Transaction{}, nil
}

func ListUserTransactionsById(id string) []model.Transaction {

	return []model.Transaction{}
}

func UpdateTransaction() model.Transaction {
	return model.Transaction{}
}
