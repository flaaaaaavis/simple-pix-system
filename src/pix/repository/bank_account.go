package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"projeto.com/src/pix/model"
	"projeto.com/src/pix/service"
)

type bankAccountRepository struct {
	gormConnection *gorm.DB
}

func (b bankAccountRepository) CreateBankAccount(newBankAccount *model.BankAccount) (*model.BankAccount, error) {
	err := b.gormConnection.Create(newBankAccount)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new bank account: %v", err.Error)

		return nil, err.Error
	}

	return newBankAccount, nil
}

func (b bankAccountRepository) GetBankAccountById(id string) (*model.BankAccount, error) {
	bankAccount := &model.BankAccount{}
	condition := fmt.Sprintf("id=%v", id)

	result := b.gormConnection.First(model.BankAccount{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting BankAccount from id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err.Error != nil {
		fmt.Sprintf("Error when getting bankAccount: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, bankAccount)
	if err.Error != nil {
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
