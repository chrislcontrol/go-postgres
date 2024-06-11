package factory

import (
	"github.com/chrislcontrol/go-postgres/internal/controller"
	"github.com/chrislcontrol/go-postgres/internal/db"
	"github.com/chrislcontrol/go-postgres/internal/handler"
	"github.com/chrislcontrol/go-postgres/internal/repository"
	"github.com/chrislcontrol/go-postgres/internal/usecase"
)

func BuildProductHandler() *handler.ProductHandler {
	// Repository
	repo := repository.NewProductRepository(db.GetDBSession())

	// UseCases
	createProductUseCase := usecase.NewCreateProductUseCase(repo)
	listProducts := usecase.NewListProducts(repo)

	// Controller
	ctrl := controller.NewProductController(createProductUseCase, listProducts)

	return handler.NewProductHandler(&ctrl)
}