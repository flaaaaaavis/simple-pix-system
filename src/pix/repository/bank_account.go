package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"projeto.com/src/config"
	"projeto.com/src/pix/model"
)

type BankAccountRepository struct {
	gormConnection config.Gorm
}

func (b *BankAccountRepository) CreateBankAccount(account *model.BankAccount) (*model.BankAccount, error) {
	db := b.gormConnection.Db()

	id := uuid.New().String()

	newBankAccount := &model.BankAccount{
		ID:            id,
		BankCode:      account.BankCode,
		BankName:      account.BankName,
		BankBranch:    account.BankBranch,
		AccountNumber: account.AccountNumber,
	}

	err := db.Create(newBankAccount)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new bank account: %v", err.Error)

		return nil, err.Error
	}

	return newBankAccount, nil
}

func (b *BankAccountRepository) GetBankAccountById(id string) (*model.BankAccount, error) {
	db := b.gormConnection.Db()

	bankAccount := &model.BankAccount{}

	condition := fmt.Sprintf("id=%v", id)

	result := db.First(model.BankAccount{}, condition)
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
