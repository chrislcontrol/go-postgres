package handler

import (
	"net/http"
	"github.com/chrislcontrol/go-postgres/internal/controller"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productController *controller.ProductController
}

func NewProductHandler(productController *controller.ProductController) *ProductHandler {
	return &ProductHandler{
		productController: productController,
	}
}

func (ph ProductHandler) HandleCreate(ctx *gin.Context) {
	var input struct {
		Name  string  `form:"name" binding:"required"`
		Price float64 `form:"price" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	response := ph.productController.CreateProduct(input.Name, input.Price)

	ctx.JSON(http.StatusCreated, response)
}


func (ph ProductHandler) HandleList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ph.productController.ListAll())
}
