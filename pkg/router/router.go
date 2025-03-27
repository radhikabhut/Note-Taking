package router

import (
	"note-taking/pkg/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(server *gin.Engine) {
	server.POST("/uploadFile", handler.UploadFile)
	server.GET("/render", handler.RenderMarkdown)
	server.GET("/list", handler.ListFiles)
}
