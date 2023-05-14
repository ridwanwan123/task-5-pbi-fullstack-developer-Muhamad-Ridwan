package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase()  {
	dsn := "host=localhost user=postgres password=ridwanwan dbname=vix_restapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	DB = db

	DB.AutoMigrate(&User{})
}
