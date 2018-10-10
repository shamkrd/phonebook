package repositories

import "github.com/shamkrd/phoneBookTest/models"

// ContactRepository is интерфейс для работы с контактами
type ContactRepository interface {
	GetList() ([]*models.Contact, error)
	Update(*models.Contact) error
	Delete(ID int64) error
	Add(contact *models.Contact) (int64, error)
}
