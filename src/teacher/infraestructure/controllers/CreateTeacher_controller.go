package controllers

import (
	"api/src/teacher/application"
	userApplication "api/src/user/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTeacherController struct {
	teacherSaver *application.AddTeacher
	notify       *userApplication.SendBroadcastNotification
}

func NewSaveTeacherController(
	useCase *application.AddTeacher,
	notify *userApplication.SendBroadcastNotification,
) *CreateTeacherController {
	return &CreateTeacherController{
		teacherSaver: useCase,
		notify:       notify,
	}
}

func (ts *CreateTeacherController) Run(c *gin.Context) {
	var body struct {
		Name       string `json:"name"`
		Asignature string `json:"asignature"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ts.teacherSaver.Execute(body.Name, body.Asignature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ts.notify != nil {
		if err := ts.notify.Execute(
			"Nuevo maestro registrado",
			"Se registró el maestro "+body.Name,
			map[string]string{
				"type":       "new_teacher",
				"name":       body.Name,
				"asignature": body.Asignature,
			},
		); err != nil {
			log.Printf("[WARN] no se pudo enviar notificación de nuevo maestro: %v", err)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Teacher added successfully"})
}
