package application

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NormalizeEmail quita espacios y pasa a minúsculas
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// IsValidEmailFormat comprueba el formato del correo
func IsValidEmailFormat(email string) bool {
	return email != "" && emailRegex.MatchString(email)
}
