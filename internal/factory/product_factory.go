package factory

import (
	"github.com/chrislcontrol/go-postgres/internal/controller"
	"github.com/chrislcontrol/go-postgres/internal/handler"
	"github.com/chrislcontrol/go-postgres/internal/repository"
	"github.com/chrislcontrol/go-postgres/internal/usecase"
	"gorm.io/gorm"
)

func BuildProductHandler(dbSession *gorm.DB) *handler.ProductHandler {
	// Repository
	repo := repository.NewProductRepository(dbSession)

	// UseCases
	createProductUseCase := usecase.NewCreateProductUseCase(repo)
	listProducts := usecase.NewListProducts(repo)

	// Controller
	ctrl := controller.NewProductController(createProductUseCase, listProducts)

	return handler.NewProductHandler(&ctrl)
}