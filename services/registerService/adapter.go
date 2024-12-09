package registerService

import (
	"abcShop/Dtos/registerDto"
	"abcShop/enums"
	"abcShop/models"
	"abcShop/pkg/abcPassword"
	"abcShop/pkg/abcTime"
	"abcShop/pkg/abcToken"
	"abcShop/repository/userRepository"

	"github.com/google/uuid"
)

type IRegisterService interface {
	Execute(request registerDto.Request) (*registerDto.Response, error)
}

type registerService struct {
	UserRepo userRepository.IUserRepository
}

func NewRegisterService(repo userRepository.IUserRepository) IRegisterService {
	return &registerService{
		UserRepo: repo,
	}
}

func (s *registerService) Execute(request registerDto.Request) (*registerDto.Response, error) {
	// pwd = abcToken.NewToken(request.Password, "http://localhost:8080")
	password := abcPassword.HashPassword(request.Password)
	userModel := models.User{
		Id:       uuid.New(),
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Role:     enums.USER_ROLE_MEMBER,
	}

	err := s.UserRepo.Create(&userModel)
	if err != nil {
		return nil, err
	}

	tokenBuilder := abcToken.NewToken(userModel.Id.String(), "http://localhost:8080")
	tokenString, err := tokenBuilder.GenerateToken(userModel.Id.String(), userModel.Email, "member")
	if err != nil {
		return nil, err
	}
	response := registerDto.Response{
		User: registerDto.UserResponse{
			Id:        userModel.Id.String(),
			Username:  request.Username,
			Email:     userModel.Email,
			CreatedAt: userModel.CreatedAt.Format(abcTime.LAYOUT_ISO),
		},
		Token: tokenString,
	}
	return &response, nil
}
