package controllers

import (
	"api/src/user/application"
	"errors"
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

// GetPasswordRequirements devuelve los requisitos de contraseña para mostrar en la UI del registro
func (rc *RegisterController) GetPasswordRequirements(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"description": "Tu contraseña debe cumplir con:",
		"rules": []string{
			"Mínimo 8 caracteres",
			"Al menos una letra y un número",
		},
		"example": "MiSegura123",
		"hint":     "Usa letras y números para mayor seguridad",
	})
}

func (rc *RegisterController) Run(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		RespondError(c, http.StatusBadRequest, MsgInvalidInput, CodeInvalidInput)
		return
	}

	user, err := rc.register.Execute(body.Email, body.Password, body.Name)
	if err != nil {
		var appErr *application.AppError
		if errors.As(err, &appErr) {
			RespondError(c, http.StatusBadRequest, appErr.Message, appErr.Code)
			return
		}
		RespondError(c, http.StatusInternalServerError, "Ha ocurrido un error. Inténtalo más tarde.", CodeInternalError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"user":    user,
		"password_requirements": gin.H{
			"description": "Para futuras referencias, tu contraseña debe tener:",
			"rules":       []string{"Mínimo 8 caracteres", "Al menos una letra y un número"},
			"example":    "MiSegura123",
			"hint":       "Combina letras y números para mayor seguridad",
		},
	})
}
