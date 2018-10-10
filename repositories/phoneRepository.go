package repositories

import "github.com/shamkrd/phoneBookTest/models"

//PhoneRepository is интерфейс для работы с телефонами
type PhoneRepository interface {
	GetList(contactID int64) ([]*models.Phone, error)
	Delete(ID int64) error
	Update(phone *models.Phone) error
	Add(phone *models.Phone) (int64, error)
}
