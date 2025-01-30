package main

import (
	"ApiHEXaGONAL/application"
	"ApiHEXaGONAL/config"
	"ApiHEXaGONAL/infrastructure"
	"ApiHEXaGONAL/infrastructure/handlers"
	"ApiHEXaGONAL/infrastructure/repositories"
	"log"
)

func main() {

	config.ConnectDatabase()

	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	userService := application.NewUserService(userRepo)
	productService := application.NewProductService(productRepo)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	r := infrastructure.SetupRouter(userHandler, productHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
