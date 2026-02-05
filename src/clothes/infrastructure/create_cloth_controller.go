// infrastructure/create_cloth_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type CreateClothController struct {
	createUseCase *application.CreateClothUseCase
}

func NewCreateClothController(createUseCase *application.CreateClothUseCase) *CreateClothController {
	return &CreateClothController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateClothController) Run(c *gin.Context) {
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