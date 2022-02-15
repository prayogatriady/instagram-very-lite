package main

import (
	"log"
	"os"
	"test-mongodb/controller"
	"test-mongodb/database"
	"test-mongodb/repository"
	"test-mongodb/routes"
	"test-mongodb/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	// to access .env file with os.Getenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get port from .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	databaseName := "mongo-golang"
	collectionName := "users"
	db := database.DBInit(databaseName, collectionName)
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, *validate)
	userController := controller.NewUserController(userService)

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router, userController)
	routes.UserRoutes(router, userController)

	log.Printf("Server running on port %s.. \n", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error when running port %s \n", port)
	}
}
