package main

import (
	"awesomeProject/controllers"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=prince password=prince dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	configController := controllers.NewConfigController(db)

	http.HandleFunc("/config/get", configController.GetConfigHandler)
	http.HandleFunc("/config/post", configController.UpdateConfigHandler)

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
