package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shamkrd/phoneBookTest/models"
)

//PhoneRepository структура для взаимодействия с телефонами
type PhoneRepository struct {
	Database *sql.DB
}

//GetList is получение списка телефонов по id клиента
func (pr *PhoneRepository) GetList(contactID int64) ([]*models.Phone, error) {
	println(pr.Database)
	rows, err := pr.Database.Query("SELECT * FROM phonebook.phones WHERE contact_id = ?;", contactID)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	println("OK")
	defer rows.Close()
	phones := []*models.Phone{}

	for rows.Next() {
		p := models.Phone{}
		err := rows.Scan(&p.ID, &p.Phone, &p.ContactID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		phones = append(phones, &p)
	}
	return phones, nil
}

//Delete is удаление телефона
func (pr *PhoneRepository) Delete(ID int64) error {
	_, err := pr.Database.Exec("DELETE FROM phonebook.phones  WHERE id = ?;",
		ID)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Update is удаление телефона
func (pr *PhoneRepository) Update(phone *models.Phone) error {
	_, err := pr.Database.Exec("UPDATE phonebook.phones SET phone = ?  WHERE id = ?;",
		phone.Phone, phone.ID)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Add is добавленеие телефона
func (pr *PhoneRepository) Add(phone *models.Phone) (int64, error) {

	result, err := pr.Database.Exec("insert into phonebook.phones (phone, contact_id) values (?,?)", phone.Phone, phone.ContactID)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		println(err)
		return 0, err
	}

	return id, nil
}
