package core

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ContextKey tipo para las claves del contexto (evita colisiones)
type ContextKey string

const (
	// ClaimsContextKey es la clave donde se guardan los claims del JWT en el contexto
	ClaimsContextKey ContextKey = "jwt_claims"
)

// AuthMiddleware retorna un middleware que valida el JWT y guarda los claims en el contexto
func AuthMiddleware(jwtRepo *JWTRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "falta el header Authorization"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "formato de Authorization inválido. Usa: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := jwtRepo.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido o expirado"})
			c.Abort()
			return
		}

		c.Set(string(ClaimsContextKey), claims)
		c.Next()
	}
}
