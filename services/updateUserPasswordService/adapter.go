package updateUserPasswordService

import (
	"abcShop/models"
	"abcShop/pkg/abcPassword"
	"abcShop/pkg/abcToken"
	"abcShop/repository/userRepository"
	"time"
)

type IUpdateUserPasswordService interface {
	Execute(request Request) (*Response, error)
}

type updateUserPasswordService struct {
	repo userRepository.IUserRepository
}

func NewUpdateUserPasswordService(repo userRepository.IUserRepository) IUpdateUserPasswordService {
	return &updateUserPasswordService{
		repo: repo,
	}
}

func (s *updateUserPasswordService) Execute(request Request) (*Response, error) {
	userModel, err := s.repo.FindOne(&models.User{Id: request.Id})
	if err != nil {
		return nil, err
	}

	// hash password
	hashPassword := abcPassword.HashPassword(request.NewPassword)
	userModel.HashPassword = hashPassword

	// update user in db
	err = s.repo.Update(userModel)
	if err != nil {
		return nil, err
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
			CreatedAt: userModel.CreatedAt.Format(time.ANSIC),
		},
		Token: tokenString,
	}
	return &response, nil
}
