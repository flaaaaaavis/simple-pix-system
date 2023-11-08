package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
)

type ContactRepository struct {
	gormConnection *gorm.DB
}

func (c *ContactRepository) CreateContact(contact *model.Contact) (*model.Contact, error) {
	id := uuid.New().String()

	newContact := &model.Contact{
		ID:          id,
		PhoneNumber: contact.PhoneNumber,
		Email:       contact.Email,
	}

	err := c.gormConnection.Create(newContact)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new contact: %v", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func (c *ContactRepository) GetContactById(id string) (*model.Contact, error) {
	contact := &model.Contact{}

	condition := fmt.Sprintf("id=%v", id)

	result := c.gormConnection.First(model.Contact{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting contact: %v", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting contact: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, contact)
	if err != nil {
		fmt.Sprintf("Error when getting contact: %v", err.Error())

		return nil, errors.New(err.Error())
	}

	return contact, nil
}

func (c *ContactRepository) UpdateContactById(newContact *model.Contact) (*model.Contact, error) {
	contact := &model.Contact{
		PhoneNumber: newContact.PhoneNumber,
		Email:       newContact.Email,
	}

	err := c.gormConnection.Model(contact).Where("id IN (?)", newContact.ID).Updates(contact)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %v", err.Error)

		return nil, err.Error
	}

	return contact, nil
}

func NewContact(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		gormConnection: db,
	}
}
