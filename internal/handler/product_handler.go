package handler

import (
	"net/http"

	"github.com/chrislcontrol/go-postgres/internal/controller"
	"github.com/chrislcontrol/go-postgres/internal/db"
	"github.com/chrislcontrol/go-postgres/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
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
	var query struct {
		Limit uint `form:"limit"`
	}
	ctx.ShouldBindQuery(&query)

	if query.Limit == 0 {
		query.Limit = 10
	}

	paginator, err := pagination.New(pagination.Options{
		GinContext:    ctx,
		DB:            db.GetDBSession(),
		Model:         &entity.Product{},
		Limit:         uint(query.Limit),
		DefaultCursor: nil,
	})
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, ph.productController.ListAll(paginator))
}
