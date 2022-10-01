package currency

import (
	"time"

	"github.com/denizedizcan/Golang-Curmin/common/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyShow struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Currrency struct {
	Currency CurrencyShow
}
type CurrencyData struct {
	Currencies []CurrencyShow
}

func GetCurrencies(db *gorm.DB) ([]CurrencyShow, error) {
	var currencies []CurrencyShow

	if result := db.Model(models.Currency{}).Find(&currencies); result.Error != nil {
		return []CurrencyShow{}, result.Error
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

func GetCurrencyByBase(base string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if result := db.Model(models.Currency{}).Preload(clause.Associations).Where("code = ?", base).First(&currency); result.Error != nil {
		return models.Currency{}, result.Error
	}
	return currency, nil
}
func GetCurrencyByBaseAndTarget(base, target string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency

	result := db.Preload("CurrencyData", "base = ?", target).Preload(clause.Associations).Where("currencies.code = ?", base).First(&currency)
	if result.Error != nil {
		return models.Currency{}, result.Error
	}

	return currency, nil
}

func GetCurrencyBytime(base, target string, start_date, end_date time.Time, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency

	result := db.Preload("CurrencyData", "base = ? AND date <= ? AND date >= ?", target, end_date, start_date).Preload(clause.Associations).Where("currencies.code = ?", base).First(&currency)
	if result.Error != nil {
		return models.Currency{}, result.Error
	}

	return currency, nil
}
