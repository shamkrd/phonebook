package startup

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/shamkrd/phoneBookTest/controllers"
	"github.com/shamkrd/phoneBookTest/models"
	mysql "github.com/shamkrd/phoneBookTest/repositories/mySql"
	"github.com/tkanos/gonfig"
)

//Programm is инициалищация компонентов перед запуском приложения
type Programm struct {
}

//Run is запускает приложение
func (p *Programm) Run() {

	configuration := models.Configuration{}
	err := gonfig.GetConf("config.development.json", &configuration)

	if err != nil {
		println(err)
		return
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	dbConfig := configuration.GetDbConnectionString()

	db, err := sql.Open("mysql", dbConfig)
	defer db.Close()
	if err != nil {
		log.Println(err)
	}

	contactRepository := mysql.ContactRepository{Database: db}
	phoneRepository := mysql.PhoneRepository{Database: db}
	contactController := controllers.ContactController{ContactRepository: &contactRepository}
	phoneController := controllers.PhoneController{PhoneRepository: &phoneRepository}

	http.HandleFunc("/", contactController.Index)
	http.HandleFunc("/createContact", contactController.CreateContact)
	http.HandleFunc("/updateContact", contactController.UpdateContact)
	http.HandleFunc("/deleteContact", contactController.DeleteContact)
	http.HandleFunc("/deletePhone", phoneController.DeletePhone)
	http.HandleFunc("/getPhones", phoneController.GetPhones)
	http.HandleFunc("/addPhone", phoneController.AddPhone)
	http.HandleFunc("/updatePhone", phoneController.UpdatePhone)

	fmt.Println("Server is listerning...")

	http.ListenAndServe(configuration.Port, nil)

}
