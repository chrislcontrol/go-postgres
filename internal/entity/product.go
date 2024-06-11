package entity

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name" gorm:"size:50"`
	Price     float64        `json:"price"`
}
