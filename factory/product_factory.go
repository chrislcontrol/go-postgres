package factory

import (
	"github.com/chrislcontrol/go-postgres/controller"
	"github.com/chrislcontrol/go-postgres/handler"
	"github.com/chrislcontrol/go-postgres/repository"
	"github.com/chrislcontrol/go-postgres/usecase"
	"gorm.io/gorm"
)

func BuildProductHandler(dbSession *gorm.DB) *handler.ProductHandler {
	// Repository
	repo := repository.NewProductRepository(dbSession)

	// UseCase
	useCase := usecase.NewCreateProductUseCase(&repo)

	// Controller
	ctrl := controller.NewProductController(&useCase)

	// Handler
	h := handler.NewProductHandler(&ctrl)

	return &h
}