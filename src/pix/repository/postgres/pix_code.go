package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mentoria/src/pix/model/postgres"
	"mentoria/src/pix/service"
)

type pixCodeRepository struct {
	gormConnection *gorm.DB
}

func (pc pixCodeRepository) CreatePixCode(newPixCode *postgres.PixCode) (*postgres.PixCode, error) {
	err := pc.gormConnection.Create(newPixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new pix code: %v", err.Error)

		return nil, err.Error
	}

	return newPixCode, nil
}

func (pc pixCodeRepository) GetPixCodeByPixId(id string) (*postgres.PixCode, error) {
	PixCode := &postgres.PixCode{}
	condition := fmt.Sprintf("pix_id=%v", id)

	result := pc.gormConnection.First(postgres.PixCode{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting PixCode from pix_id: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, PixCode)
	if err != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return PixCode, nil
}

func (pc pixCodeRepository) GetPixCodeByCode(code string) (*postgres.PixCode, error) {
	PixCode := &postgres.PixCode{}

	condition := fmt.Sprintf("code=%v", code)

	result := pc.gormConnection.First(postgres.PixCode{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting PixCode from code: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, PixCode)
	if err != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return PixCode, nil
}

func (pc pixCodeRepository) UpdatePixCode(newPixCode postgres.PixCode) (*postgres.PixCode, error) {
	err := pc.gormConnection.Model(newPixCode).Where("id IN (?)", newPixCode.ID).Updates(newPixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %v", err.Error)

		return nil, err.Error
	}

	return &newPixCode, nil
}

func (pc pixCodeRepository) DeletePixCode(code string) error {
	PixCode := &postgres.PixCode{}
	condition := fmt.Sprintf("code=%v", code)

	err := pc.gormConnection.Where(condition).Delete(PixCode)
	if err.Error != nil {
		fmt.Sprintf("Error when getting PixCode: %v", err)

		return err.Error
	}

	return nil
}

func NewPixCode(db *gorm.DB) service.PixCodeRepo {
	return &pixCodeRepository{
		gormConnection: db,
	}
}
