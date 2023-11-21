package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mentoria/src/pix/model/postgres"
	"mentoria/src/pix/service"
)

type bankAccountRepository struct {
	gormConnection *gorm.DB
}

func (b bankAccountRepository) CreateBankAccount(newBankAccount *postgres.BankAccount) (*postgres.BankAccount, error) {
	err := b.gormConnection.Create(newBankAccount)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new bank account: %v", err.Error)

		return nil, err.Error
	}

	return newBankAccount, nil
}

func (b bankAccountRepository) GetBankAccountById(id string) (*postgres.BankAccount, error) {
	bankAccount := &postgres.BankAccount{}
	condition := fmt.Sprintf("id=%v", id)

	result := b.gormConnection.First(postgres.BankAccount{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting BankAccount from id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting bankAccount: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, bankAccount)
	if err != nil {
		fmt.Sprintf("Error when getting bankAccount: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return bankAccount, nil
}

func NewBankAccount(db *gorm.DB) service.BankAccountRepo {
	return &bankAccountRepository{
		gormConnection: db,
	}
}
