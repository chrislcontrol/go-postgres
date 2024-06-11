package router

import (
	"net/http"

	"github.com/chrislcontrol/go-postgres/factory"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const apiV1 = "/api/v1"


func Route(server *gin.Engine, dbSession *gorm.DB) {
	// Ping
	server.GET(apiV1+"/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Products
	server.POST(apiV1+"/products", factory.BuildProductHandler(dbSession).HandleCreate)
}