package application

import "unicode"

const minPasswordLength = 8

// ValidatePasswordStrength exige mínimo 8 caracteres, al menos una letra y un número
func ValidatePasswordStrength(password string) bool {
	if len(password) < minPasswordLength {
		return false
	}
	hasLetter := false
	hasNumber := false
	for _, r := range password {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
		if hasLetter && hasNumber {
			return true
		}
	}
	return hasLetter && hasNumber
}

// MinPasswordLength para exponer a la API de requisitos
func MinPasswordLength() int {
	return minPasswordLength
}
