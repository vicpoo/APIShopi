// infrastructure/get_cloth_by_id_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type GetClothByIDController struct {
	getByIDUseCase *application.GetClothByIDUseCase
}

func NewGetClothByIDController(getByIDUseCase *application.GetClothByIDUseCase) *GetClothByIDController {
	return &GetClothByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

func (ctrl *GetClothByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	cloth, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Prenda no encontrada",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cloth)
}