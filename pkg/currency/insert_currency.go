package currency

import (
	"github.com/denizedizcan/Golang-Curmin/common/models"
	"gorm.io/gorm"
)

func InsertCurrency(db *gorm.DB, c []models.Currency) error {
	if result := db.Create(&c); result.Error != nil {
		return result.Error
	}
	return nil
}
