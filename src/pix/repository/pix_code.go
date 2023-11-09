package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projeto.com/src/pix/model"
)

type PixCodeRepository struct {
	gormConnection *gorm.DB
}

func (pc *PixCodeRepository) CreatePixCode(pixCode *model.PixCode) (*model.PixCode, error) {
	db := pc.gormConnection

	id := uuid.New().String()

	newPixCode := &model.PixCode{
		ID:    id,
		PixID: pixCode.PixID,
		Type:  pixCode.Type,
		Code:  pixCode.Code,
	}

	err := db.Create(newPixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new pix code: %v", err.Error)

		return nil, err.Error
	}

	return newPixCode, nil
}

func (pc *PixCodeRepository) GetPixCodeByPixId(id string) (*model.PixCode, error) {
	db := pc.gormConnection

	PixCode := &model.PixCode{}

	condition := fmt.Sprintf("pix_id=%v", id)

	result := db.First(model.PixCode{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting PixCode from pix_id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, PixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return PixCode, nil
}

func (pc *PixCodeRepository) GetPixCodeByCode(code string) (*model.PixCode, error) {
	db := pc.gormConnection

	PixCode := &model.PixCode{}

	condition := fmt.Sprintf("code=%v", code)

	result := db.First(model.PixCode{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting PixCode from code: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, PixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return PixCode, nil
}

func (pc *PixCodeRepository) UpdatePixCode(newPixCode model.PixCode) (*model.PixCode, error) {
	db := pc.gormConnection

	pixCode := &model.PixCode{
		Code: newPixCode.Code,
	}

	err := db.Model(newPixCode).Where("id IN (?)", newPixCode.ID).Updates(pixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %v", err.Error)

		return nil, err.Error
	}

	return &newPixCode, nil
}

func (pc *PixCodeRepository) DeletePixCode(code string) error {
	db := pc.gormConnection

	PixCode := &model.PixCode{}

	condition := fmt.Sprintf("code=%v", code)

	err := db.Where(condition).Delete(PixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err)

		return err.Error
	}

	return nil
}
