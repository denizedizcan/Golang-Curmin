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

func GetCurrencyId(base string, db *gorm.DB) (uint64, error) {
	var currency models.Currency
	if result := db.Model(models.Currency{}).Preload(clause.Associations).Where("code = ?", base).First(&currency); result.Error != nil {
		return 0, result.Error
	}
	return currency.ID, nil
}
