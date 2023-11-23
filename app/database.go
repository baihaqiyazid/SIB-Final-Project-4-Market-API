package app

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "marketplace"
	password = ""
	dbPort   = "5432"
	dbName   = "marketplace"
	db       *gorm.DB
	err      error
	// dns      = "postgres://postgres:postgres@localhost:5432/goreactmovies?sslmode=disable"
)

func StartDB()  {
	config := "host=localhost user=postgres password='' dbname=marketplace port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil{
		log.Fatal("error connecting to database:", err)
	}

	// db.Debug().AutoMigrate( models.User{}, models.Product{}, models.Category{}, models.TransactionHistory{})
}

func GetDB() *gorm.DB{
	return db
}