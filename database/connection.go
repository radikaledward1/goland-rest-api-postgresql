package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=secretpassword dbname=go_crud port=5432"
var DB *gorm.DB

func DbConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Connection Success wuth PostgreSQL")
	}
}
