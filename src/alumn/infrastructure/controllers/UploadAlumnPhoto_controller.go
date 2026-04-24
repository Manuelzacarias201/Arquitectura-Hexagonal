package controllers

import (
	"api/src/alumn/application"
	"api/src/core"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadAlumnPhotoController struct {
	updatePhoto *application.UpdateAlumnPhoto
}

func NewUploadAlumnPhotoController(updatePhoto *application.UpdateAlumnPhoto) *UploadAlumnPhotoController {
	return &UploadAlumnPhotoController{updatePhoto: updatePhoto}
}

func (uc *UploadAlumnPhotoController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumn ID"})
		return
	}

	photoPath, err := core.SaveUploadedFile(c, "photo", "uploads/alumns")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.updatePhoto.Execute(id, photoPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Photo updated successfully",
		"photo_path": photoPath,
	})
}
