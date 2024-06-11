package usecase

import (
	"github.com/chrislcontrol/go-postgres/internal/entity"
	"github.com/chrislcontrol/go-postgres/internal/repository"
)

type ListProducts struct {
	repo *repository.ProductRepository
}

func NewListProducts(repo *repository.ProductRepository) *ListProducts {
	return &ListProducts{repo: repo}
}

func (lp ListProducts) Execute() []entity.Product {
	return lp.repo.ListAll()
}
