package main

import (
	"log"
	"os"
	"test-mongodb/controller"
	"test-mongodb/database"
	"test-mongodb/middleware"
	"test-mongodb/repository"
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

	r := gin.New()

	r.Use(gin.Logger())

	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)
	r.POST("/logout", userController.Logout)

	r.Use(middleware.AuthMiddleware)                 // middleware, all the endpoints below access this
	r.GET("/users", userController.FindAllUser)      // get all users, user_type must be admin
	r.GET("/profile", userController.Profile)        // get profile detail, included feeds
	r.PUT("/edit", userController.EditProfile)       // edit profile
	r.DELETE("/delete", userController.DeleteUser)   // delete/deactivate profile
	r.POST("/createfeed", userController.CreateFeed) // post feed

	log.Printf("Server running on Port %s.. \n", port)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error when running port %s \n", port)
	}
}
