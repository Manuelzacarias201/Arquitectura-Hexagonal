package core

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveBase64Image guarda una imagen base64 (con o sin prefijo data:) y retorna la ruta pública.
func SaveBase64Image(base64Image, destinationFolder string) (string, error) {
	if strings.TrimSpace(base64Image) == "" {
		return "", nil
	}

	rawData := strings.TrimSpace(base64Image)
	extension := ".jpg"

	if strings.HasPrefix(rawData, "data:") {
		parts := strings.SplitN(rawData, ",", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("formato base64 inválido")
		}
		meta := parts[0]
		rawData = parts[1]
		switch {
		case strings.Contains(meta, "image/png"):
			extension = ".png"
		case strings.Contains(meta, "image/webp"):
			extension = ".webp"
		default:
			extension = ".jpg"
		}
	}

	decoded, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", fmt.Errorf("no se pudo decodificar la imagen base64: %w", err)
	}

	if err := os.MkdirAll(destinationFolder, 0o755); err != nil {
		return "", fmt.Errorf("error creando carpeta de uploads: %w", err)
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	fullPath := filepath.Join(destinationFolder, fileName)
	if err := os.WriteFile(fullPath, decoded, 0o644); err != nil {
		return "", fmt.Errorf("error guardando imagen: %w", err)
	}

	publicPath := filepath.ToSlash(fullPath)
	if !strings.HasPrefix(publicPath, "/") {
		publicPath = "/" + publicPath
	}
	return publicPath, nil
}
