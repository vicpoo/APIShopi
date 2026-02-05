// routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	// Inicializar dependencias
	registerController, loginController, createController, updateController,
		deleteController, getByIDController, getAllController := InitUserDependencies()

	// Versión 1 de la API
	v1 := router.Group("/api/")
	{
		// Rutas de usuarios
		users := v1.Group("/users")
		{
			// Autenticación
			users.POST("/register", registerController.Run)
			users.POST("/login", loginController.Run)

			users.GET("/", getAllController.Run)
			users.GET("/:id", getByIDController.Run)
			users.POST("/", createController.Run)
			users.PUT("/:id", updateController.Run)
			users.DELETE("/:id", deleteController.Run)
		}
	}
}