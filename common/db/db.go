package db

import (
	"fmt"
	"log"

	"github.com/denizedizcan/Golang-Curmin/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init the db
func Init(user, password, host, port, db_name string) *gorm.DB {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dbUrl := fmt.Sprintf(dsn, host, user, password, db_name, port)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", postgres.Open(dbUrl))
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Currency{}, &models.CurrencyData{})
	return db
}
