package currency

import (
	"github.com/denizedizcan/Golang-Curmin/common/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetCurrencies(db *gorm.DB) ([]models.Currency, error) {
	var currencies []models.Currency

	if result := db.Model(models.Currency{}).Preload(clause.Associations).Find(&currencies); result.Error != nil {
		return []models.Currency{}, result.Error
	}
	return currencies, nil
}
