package core

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context, fieldName, destinationFolder string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", fmt.Errorf("archivo '%s' requerido", fieldName)
	}
	return saveMultipartFile(c, file, destinationFolder)
}

func saveMultipartFile(c *gin.Context, file *multipart.FileHeader, destinationFolder string) (string, error) {
	if err := os.MkdirAll(destinationFolder, 0o755); err != nil {
		return "", fmt.Errorf("error creando carpeta de uploads: %w", err)
	}

	extension := strings.ToLower(filepath.Ext(file.Filename))
	if extension == "" {
		extension = ".jpg"
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	relativePath := filepath.ToSlash(filepath.Join(destinationFolder, filename))
	if err := c.SaveUploadedFile(file, relativePath); err != nil {
		return "", fmt.Errorf("error guardando archivo: %w", err)
	}

	if !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}
	return relativePath, nil
}
