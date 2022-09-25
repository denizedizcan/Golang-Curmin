package currency

import (
	"fmt"
	"time"

	"github.com/denizedizcan/Golang-Curmin/common/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyBaseAndTarget struct {
	Code       string    `json:"code"`
	Date       time.Time `json:"date"`
	Currencies []Currency
}
type Currency struct {
	Base string  `json:"base"`
	Rate float64 `json:"rate"`
}

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

func GetCurrencyByBase(base string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	if result := db.Model(models.Currency{}).Preload(clause.Associations).Where("code = ?", base).First(&currency); result.Error != nil {
		return models.Currency{}, result.Error
	}
	return currency, nil
}
func GetCurrencyByBaseAndTarget(base, target string, db *gorm.DB) (models.Currency, error) {
	var currency models.Currency
	//result := db.Raw("SELECT * FROM currencies c INNER JOIN currency_data p ON p.currency_id = c.id WHERE c.code  = ? AND p.base  = ?;", base, target).Scan(&currency)
	//result := db.Model(models.Currency{}).Preload(clause.Associations).Joins("currency_data", db.Where(&models.CurrencyData{Base: target})).Where("code = ?", base).Find(&currency)

	//query := db.Table("currency_data").Select("currency_data.base as base").Joins("left join currencies currencies on currency_data.currency_id = currencies.id").Where("currencies.code = ?", base)
	//db.Model(&models.Currency{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)

	//result := db.Joins("JOIN currency_data ON currency_data.currency_id = currencies.id AND currency_data.base = ? AND currencies.code = ?", target, base).Find(&currency)
	//result := db.Preload("currency_data", "base = ?", target).First(&currency, "code = ?", base)
	//result := db.Model(models.Currency{}).Preload(clause.Associations).Where("code = ?", base).First(&currency, models.CurrencyData{Base: target})
	//result := db.Preload("Currency.CurrencyData").Preload(clause.Associations).Where("currencies.code = ? AND currency_data.base = ?", base, target).First(&currency)

	result := db.Preload("CurrencyData", "base = ?", target).Preload(clause.Associations).Where("currencies.code = ?", base).First(&currency)
	fmt.Println(currency)
	if result.Error != nil {
		return models.Currency{}, result.Error
	}

	return currency, nil
}
