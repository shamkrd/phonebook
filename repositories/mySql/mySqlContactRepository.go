package mysql

import (
	"database/sql"
	"log"

	"github.com/shamkrd/phoneBookTest/models"
)

//ContactRepository is
type ContactRepository struct {
	Database *sql.DB
}

//GetList is получение списка пользователей из базы данных mySql
func (r *ContactRepository) GetList() ([]*models.Contact, error) {

	rows, err := r.Database.Query("SELECT * FROM phonebook.contacts;")

	if err != nil {
		log.Printf("error:%v", err)
		return nil, err
	}

	defer rows.Close()

	contacts := []*models.Contact{}

	for rows.Next() {
		c := models.Contact{}
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Printf("error2:%v", err)
			return nil, err
		}
		contacts = append(contacts, &c)
	}
	return contacts, nil
}

//Add is функция добавления нового контакта
func (r *ContactRepository) Add(contact *models.Contact) (int64, error) {
	result, err := r.Database.Exec("insert into phonebook.contacts (name) values (?)", contact.Name)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		println(err)
		return 0, err
	}

	return id, err
}

//Update is функция редактирования контакта
func (r *ContactRepository) Update(contact *models.Contact) error {
	_, err := r.Database.Exec("UPDATE phonebook.contacts SET name = ? WHERE id = ?;",
		contact.Name, contact.ID)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Delete is функция удаления контакта
func (r *ContactRepository) Delete(id int64) error {
	_, err := r.Database.Exec("DELETE FROM phonebook.contacts  WHERE id = ?;",
		id)

	if err != nil {
		log.Println(err)
	}
	return err
}
