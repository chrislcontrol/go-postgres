package controller

import (
	"github.com/chrislcontrol/go-postgres/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase *usecase.CreateProductUseCase
	listProducts         *usecase.ListProducts
}

func NewProductController(
	createProductUseCase *usecase.CreateProductUseCase,
	listProducts *usecase.ListProducts,
) ProductController {
	return ProductController{
		createProductUseCase: createProductUseCase,
		listProducts: listProducts,
	}
}

func (pc ProductController) CreateProduct(name string, price float64) gin.H {
	product := pc.createProductUseCase.Execute(name, price)

	return gin.H{
		"id":    product.ID,
		"name":  product.Name,
		"price": product.Price,
	}
}

func (pc ProductController) ListAll() gin.H {
	products := pc.listProducts.Execute()

	return gin.H{
		"data": products,
	}
}
