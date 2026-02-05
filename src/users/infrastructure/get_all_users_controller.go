//get_all_users_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
)

type GetAllUsersController struct {
	getAllUseCase *application.GetAllUsersUseCase
}

func NewGetAllUsersController(getAllUseCase *application.GetAllUsersUseCase) *GetAllUsersController {
	return &GetAllUsersController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllUsersController) Run(c *gin.Context) {
	users, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los usuarios",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}