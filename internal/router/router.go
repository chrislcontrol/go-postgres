package router

import (
	"net/http"

	"github.com/chrislcontrol/go-postgres/internal/factory"
	"github.com/gin-gonic/gin"
)

const apiV1 = "/api/v1"

// Handlers
var productHandler = factory.BuildProductHandler()

func Route(server *gin.Engine) {
	// Ping
	server.GET(apiV1+"/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Products
	server.POST(apiV1+"/products", productHandler.HandleCreate)
	server.GET(apiV1+"/products", productHandler.HandleList)
}
