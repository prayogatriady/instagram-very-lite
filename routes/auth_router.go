package routes

import (
	"test-mongodb/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, userController controller.UserController) {
	router.POST("/signup", userController.Signup) // create account, need to login
	router.POST("/login", userController.Login)   // login with existing account, and then automatically set cookie
	router.POST("/logout", userController.Logout) // logout and delete cookie
}
