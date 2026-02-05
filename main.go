// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/apiShop/src/core"
	"github.com/vicpoo/apiShop/src/users/infrastructure"
	clothes_infrastructure "github.com/vicpoo/apiShop/src/clothes/infrastructure"
)

func main() {
	// Inicializar base de datos
	core.InitDB()

	// Configurar router de Gin
	router := gin.Default()

	// Configurar CORS
	router.Use(CORSMiddleware())

	// Configurar rutas de usuarios
	infrastructure.SetupUserRoutes(router)

	// Configurar rutas de prendas (clothes)
	clothes_infrastructure.SetupClothRoutes(router)

	// Iniciar servidor
	router.Run(":8000")
}

// Middleware de CORS básico
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}