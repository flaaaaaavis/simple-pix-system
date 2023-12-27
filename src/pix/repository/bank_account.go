package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"

	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
)

type bankAccountRepository struct {
	gormConnection *gorm.DB
}

func (b bankAccountRepository) CreateBankAccount(newBankAccount *model.BankAccount) (*model.BankAccount, error) {
	err := b.gormConnection.Create(newBankAccount)
	if err.Error != nil {
		log.Fatalf("Error when creating new bank account: %v", err.Error)

		return nil, err.Error
	}

	return newBankAccount, nil
}

func (b bankAccountRepository) GetBankAccountById(id string) (*model.BankAccount, error) {
	bankAccount := &model.BankAccount{}
	condition := fmt.Sprintf("id=%v", id)

	result := b.gormConnection.First(model.BankAccount{}, condition)
	if result.Error != nil {
		log.Fatalf("Error when getting BankAccount from id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		log.Fatalf("Error when getting bankAccount: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, bankAccount)
	if err != nil {
		log.Fatalf("Error when getting bankAccount: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return bankAccount, nil
}

func NewBankAccount(db *gorm.DB) service.BankAccountService {
	return &bankAccountRepository{
		gormConnection: db,
	}
}
