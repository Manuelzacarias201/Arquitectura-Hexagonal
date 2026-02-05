package controllers

import (
	"api/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	login *application.Login
}

func NewLoginController(login *application.Login) *LoginController {
	return &LoginController{
		login: login,
	}
}

func (lc *LoginController) Run(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	token, userResponse, err := lc.login.Execute(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   token,
		"user":    userResponse.User,
	})
}
