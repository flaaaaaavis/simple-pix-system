package repository

import (
	"fmt"
	"gorm.io/gorm"
	"mentoria/src/pix/model/postgres"
	"mentoria/src/pix/service"
)

type transactionRepository struct {
	gormConnection *gorm.DB
}

func (t transactionRepository) CreateTransaction(newTransaction *postgres.Transaction) (*postgres.Transaction, error) {
	err := t.gormConnection.Create(newTransaction)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new transaction: %v", err.Error)

		return nil, err.Error
	}

	return newTransaction, nil
}

func (t transactionRepository) ListUserTransactionsById(id string) ([]postgres.Transaction, error) {
	var transactions []postgres.Transaction

	condition := fmt.Sprintf("sender_id=%v OR receiver_id=%v", id, id)

	result := t.gormConnection.Where(condition).Find(&transactions)
	if result.Error != nil {
		fmt.Sprintf("Error when listing transactions: %s", result.Error)
		return nil, result.Error
	}

	return transactions, nil
}

func (t transactionRepository) UpdateTransactionById(transaction *postgres.Transaction) (*postgres.Transaction, error) {
	err := t.gormConnection.Model(transaction).Where("id IN (?)", transaction.ID).Updates(transaction)
	if err.Error != nil {
		fmt.Sprintf("Error when updating transaction: %s", err.Error)

		return nil, err.Error
	}

	return transaction, nil
}

func NewTransaction(db *gorm.DB) service.TransactionRepo {
	return &transactionRepository{
		gormConnection: db,
	}
}
