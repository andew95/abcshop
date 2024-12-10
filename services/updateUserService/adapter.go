package updateUserService

import (
	"abcShop/models"
	"abcShop/pkg/abcToken"
	"abcShop/repository/userRepository"
)

type IUpdateUserService interface {
	Execute(request Request) (*Response, error)
}

type updateUserService struct {
	repo userRepository.IUserRepository
}

func NewUserUpdateService(repo userRepository.IUserRepository) IUpdateUserService {
	return &updateUserService{
		repo: repo,
	}
}

func (s *updateUserService) Execute(request Request) (*Response, error) {
	userModel, err := s.repo.FindOne(&models.User{Id: request.Id})
	if err != nil {
		return nil, err
	}

	// assign request to model
	userModel = request.ToModel(userModel)

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

	response := ToResponse(userModel)
	response.Token = tokenString
	return &response, nil
}
