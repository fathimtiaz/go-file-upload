package api

import (
	"go-file-upload/app/api/handler"
	"go-file-upload/config"
	"go-file-upload/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(
	config *config.Config,
	fileSvc *service.FileSvc,
) *gin.Engine {
	router := gin.Default()

	fileHandler := handler.NewFileHandler(fileSvc)
	router.POST("/file", fileHandler.Upload)
	router.GET("/file/:id/info", fileHandler.GetInfo)
	router.GET("/file/:id/download", fileHandler.Download)

	return router
}
