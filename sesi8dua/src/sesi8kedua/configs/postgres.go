package configs

import (
	"fmt"
	"sesi8-project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 8881
	user   string = "nbcamp-user"
	pass   string = "nbcamp-pass"
	dbname string = "nbcamp-be"
)

func NewPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Person{}) {
		db.AutoMigrate(&models.Person{})
	}

	return db
}
