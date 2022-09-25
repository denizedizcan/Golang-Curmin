package models

type Currency struct {
	ID           uint64         `gorm:"primary_key;auto_increment" json:"currency_id"`
	Code         string         `gorm:"size:8;not null;unique;" json:"code"`
	Name         string         `gorm:"size:128;not null;" json:"name"`
	CurrencyData []CurrencyData `gorm:"foreignKey:Currency_ID;references:ID"`
}
