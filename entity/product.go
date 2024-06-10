package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}
