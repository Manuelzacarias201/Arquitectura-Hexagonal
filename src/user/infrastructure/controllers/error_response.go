package controllers

import (
	"github.com/gin-gonic/gin"
)

// ErrorBody es el formato estándar de error de la API.
// En Kotlin/Android puedes parsear siempre este JSON y mostrar en Toast:
//
//	val json = JSONObject(responseBody)
//	val message = json.optString("error", "Ha ocurrido un error")
//	Toast.makeText(context, message, Toast.LENGTH_LONG).show()
//
// Opcional: usar "code" para lógica (ej. if (code == "EMAIL_NOT_FOUND") ...)
type ErrorBody struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

// RespondError envía JSON con error y code. Siempre usar para que el cliente pueda mostrar error en Toast
func RespondError(c *gin.Context, status int, message, code string) {
	c.JSON(status, ErrorBody{Error: message, Code: code})
}

const (
	CodeInvalidInput   = "INVALID_INPUT"
	CodeInternalError  = "INTERNAL_ERROR"
)

// Mensaje genérico cuando el body JSON es inválido (evita exponer detalles internos en Toast)
const MsgInvalidInput = "Datos inválidos. Revisa los campos e inténtalo de nuevo."
