// infrastructure/find_cloth_by_size_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type FindClothBySizeController struct {
	findBySizeUseCase *application.FindClothBySizeUseCase
}

func NewFindClothBySizeController(findBySizeUseCase *application.FindClothBySizeUseCase) *FindClothBySizeController {
	return &FindClothBySizeController{
		findBySizeUseCase: findBySizeUseCase,
	}
}

func (ctrl *FindClothBySizeController) Run(c *gin.Context) {
	size := c.Query("size")
	if size == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El parámetro 'size' es requerido",
		})
		return
	}

	clothes, err := ctrl.findBySizeUseCase.Run(size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron buscar las prendas por talla",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clothes)
}