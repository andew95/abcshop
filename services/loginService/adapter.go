package loginService

import (
	"abcShop/models"
	"abcShop/pkg/abcPassword"
	"abcShop/pkg/abcTime"
	"abcShop/pkg/abcToken"
	"abcShop/repository/userRepository"
	"errors"
	"fmt"
)

type ILoginService interface {
	Login(request Request) (*Response, error)
}

type loginService struct {
	UserRepo userRepository.IUserRepository
}

func NewLoginService(repo userRepository.IUserRepository) ILoginService {
	return &loginService{
		UserRepo: repo,
	}
}

func (s *loginService) Login(request Request) (*Response, error) {
	userModel, err := s.UserRepo.FindOne(&models.User{Email: request.Email})
	if err != nil {
		return nil, err
	}

	fmt.Println(userModel)
	fmt.Println(request)
	// pwd := abcPassword.HashPassword(request.Password)
	// fmt.Println(pwd)
	result := abcPassword.CheckPasswordHash(request.Password, userModel.HashPassword)
	if !result {
		return nil, errors.New("password not match")
	}

	tokenBuilder := abcToken.NewToken(userModel.Id.String(), "http://localhost:8080")
	tokenString, err := tokenBuilder.GenerateToken(userModel.Id.String(), userModel.Email, "member")
	if err != nil {
		return nil, err
	}
	response := Response{
		User: UserResponse{
			Id:        userModel.Id.String(),
			Username:  userModel.Username,
			Email:     userModel.Email,
			CreatedAt: userModel.CreatedAt.Format(abcTime.LAYOUT_ISO),
		},
		Token: tokenString,
	}
	return &response, nil
}
