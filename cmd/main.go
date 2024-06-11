package main

import (
	"github.com/chrislcontrol/go-postgres/internal/db"
	"github.com/chrislcontrol/go-postgres/internal/router"
	"github.com/gin-gonic/gin"
)

const (
	host = "127.0.0.1"
	port = "8000"
)

func main() {
	server := gin.Default()
	db.StartDBSession()
	router.Route(server)
	server.Run(host + ":" + port)

}
