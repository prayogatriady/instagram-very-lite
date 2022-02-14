package controller

import (
	"context"
	"net/http"
	"test-mongodb/helper"
	"test-mongodb/model/web"
	"test-mongodb/service"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAllUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Signup(c *gin.Context)
	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
	DeleteUser(c *gin.Context)

	CreateFeed(c *gin.Context)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (uc *UserControllerImpl) FindAllUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if userTypeRequest := c.GetString("userType"); userTypeRequest != "ADMIN" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": "Admin required",
		})
		return
	}

	var users []web.UserResponse
	users = uc.UserService.FindUsers(ctx)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": users,
	})
}

func (uc *UserControllerImpl) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userRequest web.UserLoginRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	count, err := uc.UserService.CountByEmailPass(ctx, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": "Wrong email or password"},
		)
		return
	}

	user, err := uc.UserService.FindUser(ctx, userRequest.Email)

	token, _, err := helper.GenerateAllTokens(user.FullName, user.Email, user.Phone, user.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "Login",
	})
}

func (uc *UserControllerImpl) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
}

func (uc *UserControllerImpl) Signup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userRequest web.UserCreateRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	count, err := uc.UserService.CountByEmail(ctx, userRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "An error occured while checking for the email"},
		)
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "Email already used"},
		)
		return
	}

	if err := uc.UserService.CreateUser(ctx, userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "User created",
	})
}

func (uc *UserControllerImpl) Profile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	emailRequest := c.GetString("email")

	var user web.UserResponse
	user, err := uc.UserService.FindUser(ctx, emailRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": user,
	})
}

func (uc *UserControllerImpl) EditProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	emailRequest := c.GetString("email")

	var userRequest web.UserUpdateRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	if err := uc.UserService.UpdateUser(ctx, emailRequest, userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "Updated",
	})
}

func (uc *UserControllerImpl) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	emailRequest := c.GetString("email")

	count, err := uc.UserService.CountByEmail(ctx, emailRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "An error occured while checking for the email"},
		)
		return
	}

	if count == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "User does not exists"},
		)
		return
	}

	if err := uc.UserService.DeleteUser(ctx, emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "User deleted",
	})
}

func (uc *UserControllerImpl) CreateFeed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	emailRequest := c.GetString("email")

	var feedRequest web.FeedCreateRequest

	if err := c.BindJSON(&feedRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	if err := uc.UserService.CreateFeed(ctx, emailRequest, feedRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "Feed created",
	})
}
