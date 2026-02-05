// infrastructure/update_cloth_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type UpdateClothController struct {
	updateUseCase *application.UpdateClothUseCase
}

func NewUpdateClothController(updateUseCase *application.UpdateClothUseCase) *UpdateClothController {
	return &UpdateClothController{
		updateUseCase: updateUseCase,
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

	var clothRequest struct {
		Name        string   `json:"name" binding:"required"`
		Description *string  `json:"description,omitempty"`
		Size        *string  `json:"size,omitempty"`
		Price       *float64 `json:"price,omitempty"`
		Stock       *int32   `json:"stock,omitempty"`
		ImageURL    *string  `json:"image_url,omitempty"`
	}

	if err := c.ShouldBindJSON(&clothRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cloth := entities.NewClothFull(
		clothRequest.Name,
		clothRequest.Description,
		clothRequest.Size,
		clothRequest.Price,
		clothRequest.Stock,
		clothRequest.ImageURL,
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