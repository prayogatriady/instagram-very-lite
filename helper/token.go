package helper

import (
	"log"
	"test-mongodb/model/web"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAllTokens(fullName, email, phone, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &web.SignedDetails{
		FullName: fullName,
		Email:    email,
		Phone:    phone,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &web.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("SECRET_KEY"))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte("SECRET_KEY"))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
