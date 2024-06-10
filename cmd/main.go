package main

import (
	"net/http"

	"github.com/chrislcontrol/go-postgres/db"
	"github.com/gin-gonic/gin"
)

const (
	apiV1 = "/api/v1"
	port  = "8000"
)

func main() {
	server := gin.Default()
	_, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	server.GET(apiV1 + "/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.Run(":" + port)

}
