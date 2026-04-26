package controllers

import (
	"api/src/alumn/application"
	"api/src/core"
	userApplication "api/src/user/application"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddAlumnController struct { //estructura para guardar alumn
	alumnSaver *application.SaveAlumn
	notify     *userApplication.SendBroadcastNotification
}

func NewSaveAlumnController(
	useCase *application.SaveAlumn,
	notify *userApplication.SendBroadcastNotification,
) *AddAlumnController {
	return &AddAlumnController{
		alumnSaver: useCase,
		notify:     notify,
	}
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
		if as.notify != nil {
			if err := as.notify.Execute(
				"Nuevo alumno registrado",
				"Se registró el alumno "+name,
				map[string]string{
					"type":      "new_alumn",
					"name":      name,
					"matricula": matricula,
					"email":     email,
				},
			); err != nil {
				log.Printf("[WARN] no se pudo enviar notificación de nuevo alumno: %v", err)
			}
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
	if as.notify != nil {
		if err := as.notify.Execute(
			"Nuevo alumno registrado",
			"Se registró el alumno "+body.Name,
			map[string]string{
				"type":      "new_alumn",
				"name":      body.Name,
				"matricula": body.Matricula,
				"email":     body.Email,
			},
		); err != nil {
			log.Printf("[WARN] no se pudo enviar notificación de nuevo alumno: %v", err)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alumn saved successfully", "photo_path": photoPath})
}
