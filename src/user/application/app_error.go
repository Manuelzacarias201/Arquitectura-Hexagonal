package application

// AppError representa un error con código para la app (ej. Kotlin)
type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// Códigos de error para la API
const (
	CodeEmailRequired    = "EMAIL_REQUIRED"
	CodePasswordRequired = "PASSWORD_REQUIRED"
	CodeNameRequired     = "NAME_REQUIRED"
	CodeInvalidEmail     = "INVALID_EMAIL"
	CodeEmailTaken       = "EMAIL_TAKEN"
	CodePasswordTooWeak  = "PASSWORD_TOO_WEAK"
	CodeEmailNotFound    = "EMAIL_NOT_FOUND"
	CodeWrongPassword    = "WRONG_PASSWORD"
	CodeInvalidRefresh   = "INVALID_REFRESH_TOKEN"
)
