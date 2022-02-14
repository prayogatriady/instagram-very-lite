package helper

import (
	"test-mongodb/model/collection"
	"test-mongodb/model/web"
)

func ToUserReponse(user collection.User) web.UserResponse {
	userReponse := web.UserResponse{
		ID:        user.Id.Hex(),
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Feeds:     user.Feeds,
	}

	return userReponse
}
