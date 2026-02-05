package infrastructure

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileUploader struct {
	UploadDir    string
	MaxSize      int64 // tamaño máximo en bytes (ej: 5MB = 5 * 1024 * 1024)
	AllowedTypes []string
}

func NewFileUploader(uploadDir string) *FileUploader {
	// Crear directorio si no existe
	os.MkdirAll(uploadDir, 0755)
	
	return &FileUploader{
		UploadDir: uploadDir,
		MaxSize:   5 * 1024 * 1024, // 5MB por defecto
		AllowedTypes: []string{
			"image/jpeg",
			"image/jpg",
			"image/png",
			"image/gif",
			"image/webp",
		},
	}
}

func (fu *FileUploader) SaveFile(file *multipart.FileHeader, subfolder string) (string, error) {
	// Validar tamaño
	if file.Size > fu.MaxSize {
		return "", fmt.Errorf("archivo demasiado grande. Máximo: %d bytes", fu.MaxSize)
	}

	// Validar tipo de archivo
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}
	
	// Resetear el reader
	src.Seek(0, 0)
	
	mimeType := http.DetectContentType(buffer)
	if !fu.isAllowedType(mimeType) {
		return "", fmt.Errorf("tipo de archivo no permitido: %s", mimeType)
	}

	// Crear directorio específico si es necesario
	fullUploadDir := filepath.Join(fu.UploadDir, subfolder)
	if err := os.MkdirAll(fullUploadDir, 0755); err != nil {
		return "", err
	}

	// Generar nombre único
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), 
		strings.TrimSuffix(filepath.Base(file.Filename), ext), ext)
	
	// Guardar archivo
	dstPath := filepath.Join(fullUploadDir, fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	// Devolver ruta relativa
	return filepath.Join(subfolder, fileName), nil
}

func (fu *FileUploader) isAllowedType(mimeType string) bool {
	for _, allowed := range fu.AllowedTypes {
		if mimeType == allowed {
			return true
		}
	}
	return false
}

func (fu *FileUploader) DeleteFile(filePath string) error {
	fullPath := filepath.Join(fu.UploadDir, filePath)
	return os.Remove(fullPath)
}