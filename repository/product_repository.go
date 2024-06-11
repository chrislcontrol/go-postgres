package repository

import (
	"github.com/chrislcontrol/go-postgres/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) ProductRepository {
	return ProductRepository{
		conn: conn,
	}
}


func (repo ProductRepository) Create(name string, price float64) entity.Product {
	product := entity.Product{
		Name: name, 
		Price: price,
	}

	tx := repo.conn.Create(&product)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return product
}
