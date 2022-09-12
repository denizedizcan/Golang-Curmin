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
	connStr := "postgres://%s:%s@%s:%s/%s"
	dbUrl := fmt.Sprintf(connStr, user, password, host, port, db_name)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", postgres.Open(dbUrl))
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Currency{}, &models.CurrencyData{})
	return db
}
