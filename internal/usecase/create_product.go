package usecase

import (
	"github.com/chrislcontrol/go-postgres/internal/entity"
	"github.com/chrislcontrol/go-postgres/internal/repository"
)

type CreateProductUseCase struct {
	repo *repository.ProductRepository
}

func NewCreateProductUseCase(repo *repository.ProductRepository) CreateProductUseCase {
	return CreateProductUseCase{
		repo: repo,
	}
}


func (uc CreateProductUseCase)Execute(name string, price float64) entity.Product {
	return uc.repo.Create(name, price)
}
