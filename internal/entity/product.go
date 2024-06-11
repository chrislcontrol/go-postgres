package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:50"`
	Price float64 `json:"price"`
}
