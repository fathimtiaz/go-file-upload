package handler

import (
	"go-file-upload/internal/domain"
	"go-file-upload/internal/repository"
	"go-file-upload/internal/service"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	FileSvc *service.FileSvc
}

func NewFileHandler(fileSvc *service.FileSvc) *FileHandler {
	return &FileHandler{fileSvc}
}

func (h *FileHandler) Upload(c *gin.Context) {
	var file multipart.File
	var header *multipart.FileHeader
	var err error
	var result domain.File

	if file, header, err = c.Request.FormFile("file"); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result, err = h.FileSvc.SaveFile(c, file, header); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *FileHandler) GetInfo(c *gin.Context) {
	var file domain.File
	var err error
	var id int

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if file, err = h.FileSvc.GetFileInfo(c, repository.FileQuery{Id: id}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *FileHandler) Download(c *gin.Context) {
	var id int
	var err error
	var data []byte
	var fileName string

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if fileName, data, err = h.FileSvc.GetFileData(c, repository.FileQuery{Id: id}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	c.Writer.Write(data)
}
