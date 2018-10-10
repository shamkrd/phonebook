package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shamkrd/phoneBookTest/models"
	"github.com/shamkrd/phoneBookTest/repositories"
)

//PhoneController is контроллер для работы с телефонами
type PhoneController struct {
	PhoneRepository repositories.PhoneRepository
}

//AddPhone is добавление нового телефона
func (p *PhoneController) AddPhone(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var phone models.Phone
	err := decoder.Decode(&phone)
	if err != nil {
		println(err)
	}

	id, err := p.PhoneRepository.Add(&phone)

	if err != nil {
		println(err)
		return
	}

	response := struct {
		Message string
		ID      int64
	}{
		Message: "Телефон успешно добавлен",
		ID:      id,
	}

	js, err := json.Marshal(response)

	println(string(js))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//GetPhones is список телефонов
func (p *PhoneController) GetPhones(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var contact models.Contact
	err := decoder.Decode(&contact)
	if err != nil {
		println(err)
	}

	phones, err := p.PhoneRepository.GetList(contact.ID)

	println("OK")

	response := struct {
		Phones []*models.Phone
	}{
		Phones: phones,
	}

	js, err := json.Marshal(response)

	println(string(js))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//UpdatePhone is редактирование номера телефона
func (p *PhoneController) UpdatePhone(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var phone models.Phone
	err := decoder.Decode(&phone)
	if err != nil {
		println(err)
	}

	err = p.PhoneRepository.Update(&phone)

	if err != nil {
		println(err)
		return
	}

	response := struct {
		Message string
	}{
		Message: "Номер успешно изменен",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//DeletePhone is обрабатывает удаление контакта
func (p *PhoneController) DeletePhone(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var phone models.Phone
	err := decoder.Decode(&phone)
	if err != nil {
		println(err)
	}

	err = p.PhoneRepository.Delete(phone.ID)

	if err != nil {
		log.Println(err)
		return
	}
	response := struct {
		Message string
	}{
		Message: "телефон успешно удален",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
