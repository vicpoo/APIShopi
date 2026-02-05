// infrastructure/find_cloth_by_price_range_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type FindClothByPriceRangeController struct {
	findByPriceRangeUseCase *application.FindClothByPriceRangeUseCase
}

func NewFindClothByPriceRangeController(findByPriceRangeUseCase *application.FindClothByPriceRangeUseCase) *FindClothByPriceRangeController {
	return &FindClothByPriceRangeController{
		findByPriceRangeUseCase: findByPriceRangeUseCase,
	}
}

func (ctrl *FindClothByPriceRangeController) Run(c *gin.Context) {
	// Obtener parámetros de query
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	if minPriceStr == "" || maxPriceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Los parámetros 'min_price' y 'max_price' son requeridos",
		})
		return
	}

	// Convertir parámetros a float64
	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El parámetro 'min_price' debe ser un número válido",
			"error":   err.Error(),
		})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El parámetro 'max_price' debe ser un número válido",
			"error":   err.Error(),
		})
		return
	}

	// Validar que minPrice sea menor que maxPrice
	if minPrice > maxPrice {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El precio mínimo debe ser menor o igual al precio máximo",
		})
		return
	}

	clothes, err := ctrl.findByPriceRangeUseCase.Run(minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron buscar las prendas por rango de precio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clothes)
}