package db

import (
	"fmt"
	"log"
	"os"

	"github.com/denizedizcan/Golang-Curmin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init the db
func Init() *gorm.DB {
	connStr := "postgres://%s:%s@%s:%s/%s"
	dbUrl := fmt.Sprintf(connStr, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", postgres.Open(dbUrl))
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Currency{}, &models.CurrencyData{})
	return db
}
