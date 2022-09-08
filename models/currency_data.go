package models

import (
	"time"
)

// user struct fields used in db
type CurrencyData struct {
	ID   uint64    `gorm:"primary_key;auto_increment" json:"currency_id"`
	Base string    `gorm:"size:8;not null;" json:"name"`
	Rate float64   `gorm:"not null;" json:"rate"`
	Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}
