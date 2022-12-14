package models

import (
	"time"
)

type CurrencyData struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"currency_data_id"`
	Currency_ID uint64    `gorm:"not null;" json:"currency_id"`
	Base        string    `gorm:"size:8;not null;" json:"base"`
	Rate        float64   `gorm:"not null;" json:"rate"`
	Date        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}
