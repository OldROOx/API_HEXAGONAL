package infrastructure

import (
	"ApiHEXaGONAL/infrastructure/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter(userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Cambia "*" por tu dominio en producción
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/products", productHandler.GetProducts)
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	return r
}
