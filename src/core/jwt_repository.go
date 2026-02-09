package core

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTRepository struct {
	secretKey []byte
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWTRepository() (*JWTRepository, error) {
	// Cargar variables de entorno si no están cargadas
	_ = godotenv.Load()

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "default-secret-key-change-in-production" // Clave por defecto
	}

	return &JWTRepository{
		secretKey: []byte(secretKey),
	}, nil
}

// GenerateToken genera un token JWT de acceso para un usuario (24h)
func (jr *JWTRepository) GenerateToken(userID int, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jr.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken genera un token de refresco (7 días)
func (jr *JWTRepository) GenerateRefreshToken(userID int, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jr.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida un token JWT y retorna los claims
func (jr *JWTRepository) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return jr.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}
