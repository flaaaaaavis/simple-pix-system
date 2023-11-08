package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"projeto.com/src/config"
	"projeto.com/src/pix/model"
)

type PixRepository struct {
	gormConnection config.Gorm
}

func (p *PixRepository) CreatePix(pix *model.Pix) (*model.Pix, error) {
	db := p.gormConnection.Db()

	id := uuid.New().String()

	newPix := &model.Pix{
		ID:            id,
		UserID:        pix.UserID,
		BankAccountID: pix.BankAccountID,
		Balance:       pix.Balance,
	}

	err := db.Create(newPix)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new pix: %v", err.Error)

		return nil, err.Error
	}

	return newPix, nil
}

func (p *PixRepository) GetPixById(id string) (*model.Pix, error) {
	db := p.gormConnection.Db()

	Pix := &model.Pix{}

	condition := fmt.Sprintf("id=%v", id)

	result := db.First(model.Pix{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting Pix from id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err.Error != nil {
		fmt.Sprintf("Error when getting Pix: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, Pix)
	if err.Error != nil {
		fmt.Sprintf("Error when getting Pix: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return Pix, nil
}

func (p *PixRepository) UpdatePixBalance(newPix *model.Pix) (*model.Pix, error) {
	db := p.gormConnection.Db()

	pix := &model.Pix{
		Balance: newPix.Balance,
	}

	err := db.Model(pix).Where("id IN (?)", newPix.ID).Updates(pix)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %v", err.Error)

		return nil, err.Error
	}

	return pix, nil
}
