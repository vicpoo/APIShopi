//create_user_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type CreateUserController struct {
	createUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUseCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateUserController) Run(c *gin.Context) {
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

	createdUser, err := ctrl.createUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}