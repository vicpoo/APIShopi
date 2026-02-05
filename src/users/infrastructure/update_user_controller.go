//update_user_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type UpdateUserController struct {
	updateUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUseCase *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateUserController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var userRequest struct {
		Email    string  `json:"email" binding:"required,email"`
		Password string  `json:"password" binding:"required,min=6"`
		Name     *string `json:"name,omitempty"`
		Lastname *string `json:"lastname,omitempty"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	user := entities.NewUserFull(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Lastname,
	)
	user.SetIDUsuario(int32(id))

	updatedUser, err := ctrl.updateUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}