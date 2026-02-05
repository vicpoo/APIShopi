// infrastructure/get_all_clothes_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type GetAllClothesController struct {
	getAllUseCase *application.GetAllClothesUseCase
}

func NewGetAllClothesController(getAllUseCase *application.GetAllClothesUseCase) *GetAllClothesController {
	return &GetAllClothesController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllClothesController) Run(c *gin.Context) {
	clothes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las prendas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clothes)
}