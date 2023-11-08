package service

import (
	"fmt"
	"projeto.com/src/pix/model"
	"projeto.com/src/pix/repository"
)

type PixInterface interface {
	CreateBankAccount(account *model.BankAccount) (*model.BankAccount, error)
	GetBankAccountById(id string) (*model.BankAccount, error)

	CreatePix(pix *model.Pix) (*model.Pix, error)
	GetPixById(id string) (*model.Pix, error)
	UpdatePixBalance(newPix *model.Pix) (*model.Pix, error)

	CreatePixCode(pixCode *model.PixCode) (*model.PixCode, error)
	GetPixCodeByPixId(id string) (*model.PixCode, error)
	GetPixCodeByCode(code string) (*model.PixCode, error)
	DeletePixCode(code string) error
}

type pixService struct {
	bankAccountRepo repository.BankAccountRepository
	pixRepo         repository.PixRepository
	pixCodeRepo     repository.PixCodeRepository
	transactionRepo repository.TransactionRepository
}

func (p *pixService) CreateBankAccount(account *model.BankAccount) (*model.BankAccount, error) {
	newBankAccount, err := p.CreateBankAccount(account)
	if err != nil {
		fmt.Sprintf("Error when creating new bank account on pix service: %s", err)

		return nil, err
	}
	return newBankAccount, nil
}

func (p *pixService) GetBankAccountById(id string) (*model.BankAccount, error) {
	bankAccount, err := p.GetBankAccountById(id)
	if err != nil {
		fmt.Sprintf("Error when getting bank account from id on pix service: %s", err)

		return nil, err
	}
	return bankAccount, nil
}

func (p *pixService) CreatePix(pix *model.Pix) (*model.Pix, error) {
	newPix, err := p.CreatePix(pix)
	if err != nil {
		fmt.Sprintf("Error when creating new pix on pix service: %s", err)

		return nil, err
	}
	return newPix, nil
}

func (p *pixService) GetPixById(id string) (*model.Pix, error) {
	pix, err := p.GetPixById(id)
	if err != nil {
		fmt.Sprintf("Error when getting pix from id on pix service: %s", err)

		return nil, err
	}
	return pix, nil
}

func (p *pixService) UpdatePixBalance(newPix *model.Pix) (*model.Pix, error) {
	pix, err := p.UpdatePixBalance(newPix)
	if err != nil {
		fmt.Sprintf("Error when updating pix on pix service: %s", err)

		return nil, err
	}
	return pix, nil
}

func (p *pixService) CreatePixCode(pixCode *model.PixCode) (*model.PixCode, error) {
	newPixCode, err := p.CreatePixCode(pixCode)
	if err != nil {
		fmt.Sprintf("Error when creating new pix_code on pix service: %s", err)

		return nil, err
	}
	return newPixCode, nil
}

func (p *pixService) GetPixCodeByPixId(id string) (*model.PixCode, error) {
	pixCode, err := p.GetPixCodeByPixId(id)
	if err != nil {
		fmt.Sprintf("Error when getting pix_code from Pix_id on pix service: %s", err)

		return nil, err
	}
	return pixCode, nil
}

func (p *pixService) GetPixCodeByCode(code string) (*model.PixCode, error) {
	pixCode, err := p.GetPixCodeByCode(code)
	if err != nil {
		fmt.Sprintf("Error when getting pix_code from id on pix service: %s", err)

		return nil, err
	}
	return pixCode, nil
}

func (p *pixService) DeletePixCode(code string) error {
	err := p.DeletePixCode(code)
	if err != nil {
		fmt.Sprintf("Error when getting pix_code from id on pix service: %s", err)

		return err
	}
	return nil
}

func NewPixService(bankAccountRepo repository.BankAccountRepository, pixRepo repository.PixRepository, pixCodeRepo repository.PixCodeRepository, transactionRepo repository.TransactionRepository) PixInterface {
	return &pixService{
		bankAccountRepo: bankAccountRepo,
		pixRepo:         pixRepo,
		pixCodeRepo:     pixCodeRepo,
		transactionRepo: transactionRepo,
	}
}
