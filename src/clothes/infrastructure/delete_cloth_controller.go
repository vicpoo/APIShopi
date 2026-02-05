// infrastructure/delete_cloth_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/clothes/application"
)

type DeleteClothController struct {
	deleteUseCase *application.DeleteClothUseCase
}

func NewDeleteClothController(deleteUseCase *application.DeleteClothUseCase) *DeleteClothController {
	return &DeleteClothController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteClothController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la prenda",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Prenda eliminada exitosamente",
	})
}