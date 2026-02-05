package controllers

import (
	"api/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	register *application.Register
}

func NewRegisterController(register *application.Register) *RegisterController {
	return &RegisterController{
		register: register,
	}
}

func (rc *RegisterController) Run(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	err := rc.register.Execute(body.Email, body.Password, body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
	})
}
