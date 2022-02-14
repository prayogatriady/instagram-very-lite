package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "401 - UNAUTHORIZED",
			"message": "Not Authorized",
		})
		c.Abort()
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "401 - UNAUTHORIZED",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	getClaims, _ := token.Claims.(jwt.MapClaims)

	c.Set("fullName", getClaims["FullName"])
	c.Set("email", getClaims["Email"])
	c.Set("phone", getClaims["Phone"])
	c.Set("userType", getClaims["UserType"])

	c.Next()
}
