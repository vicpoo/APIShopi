package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type CreateClothController struct {
	createUseCase *application.CreateClothUseCase
	uploadDir     string
}

func NewCreateClothController(createUseCase *application.CreateClothUseCase) *CreateClothController {
	// Crear directorio de uploads si no existe
	uploadDir := "./uploads/clothes"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		fmt.Printf("Error al crear directorio de uploads: %v\n", err)
	}
	
	return &CreateClothController{
		createUseCase: createUseCase,
		uploadDir:     uploadDir,
	}
}

func (ctrl *CreateClothController) Run(c *gin.Context) {
	// Parsear form-data (multipart/form-data)
	name := c.PostForm("name")
	description := c.PostForm("description")
	size := c.PostForm("size")
	
	// Parsear precio
	var pricePtr *float64
	if priceStr := c.PostForm("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			pricePtr = &price
		}
	}
	
	// Parsear stock
	var stockPtr *int32
	if stockStr := c.PostForm("stock"); stockStr != "" {
		if stock, err := strconv.ParseInt(stockStr, 10, 32); err == nil {
			stockInt32 := int32(stock)
			stockPtr = &stockInt32
		}
	}

	// Manejar la imagen
	var imageURL *string
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		// Generar nombre único para el archivo
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
		filePath := filepath.Join(ctrl.uploadDir, fileName)
		
		// Guardar el archivo
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error al guardar la imagen",
				"error":   err.Error(),
			})
			return
		}
		
		// Crear URL relativa
		url := fmt.Sprintf("/uploads/clothes/%s", fileName)
		imageURL = &url
	}

	// Validar campos requeridos
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El nombre es requerido",
		})
		return
	}

	// Crear objeto cloth
	cloth := entities.NewClothFull(
		name,
		stringToPtr(description),
		stringToPtr(size),
		pricePtr,
		stockPtr,
		imageURL,
	)

	createdCloth, err := ctrl.createUseCase.Run(cloth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la prenda",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCloth)
}

// Función auxiliar para convertir string a *string
func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}