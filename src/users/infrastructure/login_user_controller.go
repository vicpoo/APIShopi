//login_user_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/users/application"
)

type LoginUserController struct {
	loginUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{
		loginUseCase: loginUseCase,
	}
}

func (ctrl *LoginUserController) Run(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	user, err := ctrl.loginUseCase.Run(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credenciales inválidas",
			"error":   err.Error(),
		})
		return
	}

	// No devolver la contraseña en la respuesta
	response := gin.H{
		"id_usuario": user.GetIDUsuario(),
		"email":      user.GetEmail(),
		"name":       user.GetName(),
		"lastname":   user.GetLastname(),
	}

	c.JSON(http.StatusOK, response)
}