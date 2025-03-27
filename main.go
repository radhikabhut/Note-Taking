package main

import (
	"note-taking/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	router.InitRouter(server)
	server.Run(":8080")
}
