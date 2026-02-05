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

type UpdateClothController struct {
	updateUseCase *application.UpdateClothUseCase
	uploadDir     string
}

func NewUpdateClothController(updateUseCase *application.UpdateClothUseCase) *UpdateClothController {
	uploadDir := "./uploads/clothes"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		fmt.Printf("Error al crear directorio de uploads: %v\n", err)
	}
	
	return &UpdateClothController{
		updateUseCase: updateUseCase,
		uploadDir:     uploadDir,
	}
}

func (ctrl *UpdateClothController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Parsear form-data
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

	// Manejar la imagen (opcional en actualización)
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
	} else {
		// Si no se sube nueva imagen, mantener la existente
		existingImage := c.PostForm("existing_image_url")
		if existingImage != "" {
			imageURL = &existingImage
		}
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
	cloth.SetIDCloth(int32(id))

	updatedCloth, err := ctrl.updateUseCase.Run(cloth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la prenda",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCloth)
}