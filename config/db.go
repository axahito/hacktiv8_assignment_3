package config

import (
	"fmt"
	"hacktiv8_assignment_3/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "akahito"
	dbPort   = "5432"
	dbname   = "hacktiv8_assignment3"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	} else {
		fmt.Println("successfully connected to database")
	}

	db.AutoMigrate(models.Customer{})
	db.AutoMigrate(models.Order{})
	db.AutoMigrate(models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
