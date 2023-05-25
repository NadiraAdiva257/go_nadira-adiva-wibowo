package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Electronic struct {
	gorm.Model
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	StockAmount int       `json:"stock_amount" form:"stock_amount"`
	Price       int       `json:"price" form:"price"`
	Category    uuid.UUID `json:"category" form:"category"`
	AdminID     int       `json:"admin_id" form:"admin_id"`
}

func (e *Electronic) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	e.Category = uuid

	return nil
}
