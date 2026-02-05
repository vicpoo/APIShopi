//register_user_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type RegisterUserController struct {
	registerUseCase *application.RegisterUserUseCase
}

func NewRegisterUserController(registerUseCase *application.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{
		registerUseCase: registerUseCase,
	}
}

func (ctrl *RegisterUserController) Run(c *gin.Context) {
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

	registeredUser, err := ctrl.registerUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo registrar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, registeredUser)
}