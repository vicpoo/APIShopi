package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func SetupClothRoutes(router *gin.Engine) {
	// Inicializar dependencias
	createController, updateController, deleteController,
		getByIDController, getAllController, findByNameController,
		findBySizeController, findByPriceRangeController := InitClothDependencies()

	// Servir archivos estáticos desde la carpeta uploads
	router.Static("/uploads", "./uploads")

	// Grupo de rutas para prendas
	clothRoutes := router.Group("/api/clothes")
	{
		// Rutas CRUD básicas
		clothRoutes.GET("/", getAllController.Run)
		clothRoutes.GET("/:id", getByIDController.Run)
		clothRoutes.POST("/", createController.Run)     
		clothRoutes.PUT("/:id", updateController.Run)   
		clothRoutes.DELETE("/:id", deleteController.Run)
		
		// Rutas de búsqueda
		clothRoutes.GET("/search/name", findByNameController.Run)
		clothRoutes.GET("/search/size", findBySizeController.Run)
		clothRoutes.GET("/search/price", findByPriceRangeController.Run)
	}
}