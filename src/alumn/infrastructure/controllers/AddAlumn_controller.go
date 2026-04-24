package controllers

import (
	"api/src/alumn/application"
	"api/src/core"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddAlumnController struct { //estructura para guardar alumn
	alumnSaver *application.SaveAlumn //puntero a la estructura de guardar alumn
}

func NewSaveAlumnController(useCase *application.SaveAlumn) *AddAlumnController { //constructor para la bd
	return &AddAlumnController{alumnSaver: useCase}
}

func (as *AddAlumnController) Run(c *gin.Context) { //maneja el post
	contentType := c.ContentType()
	if strings.HasPrefix(contentType, "multipart/") {
		name := c.PostForm("name")
		matricula := c.PostForm("matricula")
		email := c.PostForm("email")
		if name == "" || matricula == "" || email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name, matricula y email son requeridos"})
			return
		}

		photoPath := ""
		if _, err := c.FormFile("photo"); err == nil {
			savedURL, saveErr := core.SaveUploadedFile(c, "photo", "uploads/alumns")
			if saveErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": saveErr.Error()})
				return
			}
			photoPath = savedURL
		}

		if err := as.alumnSaver.Execute(name, matricula, email, photoPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Alumn saved successfully", "photo_path": photoPath})
		return
	}

	var body struct {
		Name      string `json:"name"`
		Matricula string `json:"matricula"`
		Email     string `json:"email"`
		PhotoPath string `json:"photo_path"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoPath := body.PhotoPath
	if strings.HasPrefix(strings.TrimSpace(body.PhotoPath), "data:image") {
		saved, err := core.SaveBase64Image(body.PhotoPath, "uploads/alumns")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		photoPath = saved
	}

	err := as.alumnSaver.Execute(body.Name, body.Matricula, body.Email, photoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alumn saved successfully", "photo_path": photoPath})
}
