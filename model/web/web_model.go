package web

import (
	"test-mongodb/model/collection"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserResponse struct {
	ID        string            `json:"id"`
	FullName  string            `json:"full_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	UserType  string            `json:"user_type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Feeds     []collection.Feed `json:"feeds"`
}

type UserCreateRequest struct {
	FullName string `json:"full_name" validate:"required,min=2,max=500"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	UserType string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	UserType string `json:"user_type"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type FeedCreateRequest struct {
	Caption string `json:"caption" validate:"required,min=1"`
}

type SignedDetails struct {
	FullName string
	Email    string
	Phone    string
	UserType string
	jwt.StandardClaims
}
