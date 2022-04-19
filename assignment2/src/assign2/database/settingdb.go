package database

import (
	"assign2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "adminaja"
	dbname   = "assign2"
	db       *gorm.DB
	err      error
)

func InitDB() *gorm.DB {
	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.Orders{}, models.Items{})

	return db

}
