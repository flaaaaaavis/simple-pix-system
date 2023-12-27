package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
)

type transactionRepository struct {
	gormConnection *gorm.DB
}

func (t transactionRepository) CreateTransaction(newTransaction *model.Transaction) (*model.Transaction, error) {
	err := t.gormConnection.Create(newTransaction)
	if err.Error != nil {
		log.Fatalf("Error when creating new transaction: %v", err.Error)

		return nil, err.Error
	}

	return newTransaction, nil
}

func (t transactionRepository) ListUserTransactionsById(id string) ([]model.Transaction, error) {
	var transactions []model.Transaction

	condition := fmt.Sprintf("sender_id=%v OR receiver_id=%v", id, id)

	result := t.gormConnection.Where(condition).Find(&transactions)
	if result.Error != nil {
		log.Fatalf("Error when listing transactions: %s", result.Error)
		return nil, result.Error
	}

	return transactions, nil
}

func NewTransaction(db *gorm.DB) service.TransactionService {
	return &transactionRepository{
		gormConnection: db,
	}
}
