package loginService_test

import (
	"abcShop/models"
	"abcShop/repository/userRepository"
	"abcShop/services/loginService"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type LoginTestSuite struct {
	suite.Suite
	UserRepo     userRepository.IUserRepository
	LoginService loginService.ILoginService
}

func (suite *LoginTestSuite) SetupTest() {
}

func (suite *LoginTestSuite) TestLogin() {
	type given struct {
		request loginService.Request
	}

	type when struct {
		userModel *models.User
		err       error
	}

	type then struct {
		response *loginService.Response
		err      error
	}

	tests := []struct {
		name     string
		request  given
		when     when
		expected then
	}{
		{
			name: "success",
			request: given{
				request: loginService.Request{
					Email:    "user@example.com",
					Password: "123456",
				},
			},
			when: when{
				userModel: &models.User{
					Email:        "user@example.com",
					HashPassword: "$2a$10$I7RG1kSI4wQQTzLfYNu0K.Y8tKfwxgQI63l9JIlefnaJAESNj/QZy",
				},
				err: nil,
			},
			expected: then{
				response: &loginService.Response{
					User: loginService.UserResponse{
						Email: "user@example.com",
					},
				},
			},
		},
		{
			name: "user not found",
			request: given{
				request: loginService.Request{
					Email:    "user@example.com",
					Password: "123456",
				},
			},
			when: when{
				userModel: &models.User{},
				err:       gorm.ErrRecordNotFound,
			},
			expected: then{
				response: nil,
				err:      gorm.ErrRecordNotFound,
			},
		},
		{
			name: "password not match",
			request: given{
				request: loginService.Request{
					Email:    "user@example.com",
					Password: "123456",
				},
			},
			when: when{
				userModel: &models.User{
					Email:        "user@example.com",
					HashPassword: "",
				},
				err: nil,
			},
			expected: then{
				response: nil,
				err:      errors.New("password not match"),
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			repo := userRepository.NewUserRepositoryMock()
			repo.On("FindOne", mock.Anything).Return(test.when.userModel, test.when.err)

			service := loginService.NewLoginService(repo)

			actualValue, actualErr := service.Login(test.request.request)
			if test.expected.err != nil {
				assert.Equal(suite.T(), test.expected.err, actualErr)

			}
			if test.expected.response != nil {
				fmt.Println(actualValue)
				assert.Equal(suite.T(), test.expected.response.User.Email, actualValue.User.Email)
			}
		})
	}
}
func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
