// infrastructure/find_cloth_by_name_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type FindClothByNameController struct {
	findByNameUseCase *application.FindClothByNameUseCase
}

func NewFindClothByNameController(findByNameUseCase *application.FindClothByNameUseCase) *FindClothByNameController {
	return &FindClothByNameController{
		findByNameUseCase: findByNameUseCase,
	}
}

func (ctrl *FindClothByNameController) Run(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El parámetro 'name' es requerido",
		})
		return
	}

	clothes, err := ctrl.findByNameUseCase.Run(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron buscar las prendas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clothes)
}