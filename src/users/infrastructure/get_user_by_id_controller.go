//get_user_by_id_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
)

type GetUserByIDController struct {
	getByIDUseCase *application.GetUserByIDUseCase
}

func NewGetUserByIDController(getByIDUseCase *application.GetUserByIDUseCase) *GetUserByIDController {
	return &GetUserByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

func (ctrl *GetUserByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	user, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Usuario no encontrado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}