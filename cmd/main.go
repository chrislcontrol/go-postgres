package main

import (
	"net/http"

	"github.com/chrislcontrol/go-postgres/db"
	"github.com/chrislcontrol/go-postgres/factory"
	"github.com/gin-gonic/gin"
)

const (
	apiV1 = "/api/v1"
	host  = "127.0.0.1"
	port  = "8000"
)

func main() {
	server := gin.Default()
	dbSession := db.ConnectDB()
	handlers := gin.New()
	handlers.Use(gin.Recovery())

	server.GET(apiV1+"/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	productHandler := factory.BuildProductHandler(dbSession)

	server.POST(apiV1+"/products", productHandler.HandleCreate)

	server.Run(host + ":" + port)

}
