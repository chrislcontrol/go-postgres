package controller

import (
	"github.com/chrislcontrol/go-postgres/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase *usecase.CreateProductUseCase
}

func NewProductController(createProductUseCase *usecase.CreateProductUseCase) ProductController {
	return ProductController{
		createProductUseCase: createProductUseCase,
	}
}

func (pc ProductController) CreateProduct(name string, price float64) gin.H {
	product := pc.createProductUseCase.Execute(name, price)

	return gin.H{
		"id": product.ID, 
		"name": product.Name, 
		"price": product.Price,
	}
}
