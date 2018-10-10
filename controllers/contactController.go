package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/shamkrd/phoneBookTest/models"
	"github.com/shamkrd/phoneBookTest/repositories"
)

//ContactController is контроллер контактов
type ContactController struct {
	ContactRepository repositories.ContactRepository
}

//Index главная страница
func (c *ContactController) Index(w http.ResponseWriter, r *http.Request) {

	contacts, err := c.ContactRepository.GetList()

	log.Printf("contacts: %v", contacts)

	if err != nil {
		println(err)
		return
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, contacts)
}

//CreateContact is обрабатывает добавление пользователя
func (c *ContactController) CreateContact(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var contact models.Contact
	err := decoder.Decode(&contact)
	if err != nil {
		println(err)
	}

	id, err := c.ContactRepository.Add(&contact)

	if err != nil {
		println(err)
		return
	}

	response := struct {
		Message string
		ID      int64
	}{
		Message: "Контакт успешно добавлен",
		ID:      id,
	}

	js, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//UpdateContact is обрабатывает изменение контакта
func (c *ContactController) UpdateContact(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var contact models.Contact
	err := decoder.Decode(&contact)
	if err != nil {
		println(err)
	}

	err = c.ContactRepository.Update(&contact)

	if err != nil {
		println(err)
		return
	}

	response := struct {
		Message string
	}{
		Message: "Контакт успешно изменен",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//DeleteContact is обрабатывает удаление контакта
func (c *ContactController) DeleteContact(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var contact models.Contact
	err := decoder.Decode(&contact)
	if err != nil {
		println(err)
	}

	err = c.ContactRepository.Delete(contact.ID)

	if err != nil {
		log.Println(err)
		return
	}
	response := struct {
		Message string
	}{
		Message: "Контакт успешно удален",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
