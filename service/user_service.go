package service

import (
	"context"
	"log"
	"strings"
	"test-mongodb/helper"
	"test-mongodb/model/collection"
	"test-mongodb/model/web"
	"test-mongodb/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	FindUser(ctx context.Context, email string) (web.UserResponse, error)
	FindUsers(ctx context.Context) []web.UserResponse
	CreateUser(ctx context.Context, request web.UserCreateRequest) error
	DeleteUser(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, email string, request web.UserUpdateRequest) error
	CountByEmail(ctx context.Context, email string) (int, error)
	CountByEmailPass(ctx context.Context, userRequest web.UserLoginRequest) (int, error)

	CreateFeed(ctx context.Context, email string, feed web.FeedCreateRequest) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       &validate,
	}
}

func (s *UserServiceImpl) FindUser(ctx context.Context, email string) (web.UserResponse, error) {
	user, err := s.UserRepository.FindUser(ctx, email)
	if err != nil {
		log.Println(err)
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserReponse(user)
	return userResponse, nil
}

func (s *UserServiceImpl) FindUsers(ctx context.Context) []web.UserResponse {
	users := s.UserRepository.FindUsers(ctx)

	var userReponses []web.UserResponse
	var userReponse web.UserResponse
	for _, user := range users {
		userReponse = helper.ToUserReponse(user)
		userReponses = append(userReponses, userReponse)
	}

	return userReponses
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, request web.UserCreateRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return err
	}

	timeNow, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
		return err
	}

	password := helper.HashPassword(request.Password)

	user := collection.User{
		FullName:  strings.ToUpper(request.FullName),
		Password:  password,
		Email:     request.Email,
		Phone:     request.Phone,
		UserType:  request.UserType,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, email string) error {
	err := s.UserRepository.DeleteUser(ctx, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, email string, request web.UserUpdateRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return err
	}

	timeNow, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
		return err
	}

	password := helper.HashPassword(request.Password)

	user := collection.User{
		FullName:  strings.ToUpper(request.FullName),
		Password:  password,
		Phone:     request.Phone,
		UserType:  request.UserType,
		UpdatedAt: timeNow,
	}

	err = s.UserRepository.UpdateUser(ctx, email, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *UserServiceImpl) CountByEmail(ctx context.Context, email string) (int, error) {
	count, err := s.UserRepository.CountByEmail(ctx, email)
	if err != nil {
		log.Println(err)
		return int(count), err
	}

	return int(count), nil
}

func (s *UserServiceImpl) CountByEmailPass(ctx context.Context, userRequest web.UserLoginRequest) (int, error) {

	userFound, _ := s.UserRepository.FindUser(ctx, userRequest.Email)

	if passwordIsValid := helper.VerifyPassword(userFound.Password, userRequest.Password); !passwordIsValid {
		return 0, nil
	}

	count, err := s.UserRepository.CountByEmailPass(ctx, userRequest.Email, userFound.Password)
	if err != nil {
		log.Println(err)
		return int(count), err
	}

	return int(count), nil
}

func (s *UserServiceImpl) CreateFeed(ctx context.Context, email string, request web.FeedCreateRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return err
	}

	timeNow, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	feed := collection.Feed{
		Caption:   request.Caption,
		CreatedAt: timeNow,
	}

	err = s.UserRepository.CreateFeed(ctx, email, feed)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
