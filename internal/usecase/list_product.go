package usecase

import (
	"github.com/chrislcontrol/go-postgres/internal/entity"
	"github.com/chrislcontrol/go-postgres/internal/repository"
	"github.com/rosberry/go-pagination"
)

type ListProducts struct {
	repo *repository.ProductRepository
}

func NewListProducts(repo *repository.ProductRepository) *ListProducts {
	return &ListProducts{repo: repo}
}

func (lp ListProducts) Execute(paginator *pagination.Paginator) []entity.Product {
	return lp.repo.ListPaginated(paginator)
}
