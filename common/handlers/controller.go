package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

// db handler
func New(db *gorm.DB) handler {
	return handler{db}
}
