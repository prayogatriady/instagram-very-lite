package routes

import (
	"test-mongodb/controller"
	"test-mongodb/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController controller.UserController) {
	router.Use(middleware.AuthMiddleware)                 // middleware, all the endpoints below access this
	router.GET("/users", userController.FindAllUser)      // get all users, user_type must be admin
	router.GET("/profile", userController.Profile)        // get profile detail, included feeds
	router.PUT("/edit", userController.EditProfile)       // edit profile
	router.DELETE("/delete", userController.DeleteUser)   // delete/deactivate profile
	router.POST("/createfeed", userController.CreateFeed) // post feed
}
