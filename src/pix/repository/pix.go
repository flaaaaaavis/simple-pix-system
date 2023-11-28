package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
)

type pixRepository struct {
	gormConnection *gorm.DB
}

func (p pixRepository) CreatePix(newPix *model.Pix) (*model.Pix, error) {
	err := p.gormConnection.Create(newPix)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new pix: %v", err.Error)

		return nil, err.Error
	}

	return newPix, nil
}

func (p pixRepository) GetPixById(id string) (*model.Pix, error) {
	Pix := &model.Pix{}
	condition := fmt.Sprintf("id=%v", id)

	result := p.gormConnection.First(model.Pix{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting Pix from id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting Pix: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, Pix)
	if err != nil {
		fmt.Sprintf("Error when getting Pix: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return Pix, nil
}

func (p pixRepository) UpdatePixBalance(newPix *model.Pix) (*model.Pix, error) {
	err := p.gormConnection.Model(newPix).Where("id IN (?)", newPix.ID).Updates(newPix)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %v", err.Error)

		return nil, err.Error
	}

	return newPix, nil
}

func NewPix(db *gorm.DB) service.PixRepo {
	return &pixRepository{
		gormConnection: db,
	}
}
