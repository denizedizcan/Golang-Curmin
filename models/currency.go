package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Currency struct {
	ID           uint64         `gorm:"primary_key;auto_increment" json:"currency_id"`
	Code         string         `gorm:"size:8;not null;" json:"code"`
	Name         string         `gorm:"size:8;not null;" json:"name"`
	CurrencyData []CurrencyData `gorm:"foreignKey:currency_id;references:currency_id"`
}

func FindAllCurrencys(db *gorm.DB) ([]Currency, error) {
	var currencies []Currency

	if result := db.Model(Currency{}).Preload(clause.Associations).Find(&currencies); result.Error != nil {
		return []Currency{}, result.Error
	}
	return currencies, nil
}
