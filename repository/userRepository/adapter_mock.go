package userRepository

import (
	"abcShop/models"

	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}
func (m *userRepositoryMock) Find() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}
func (m *userRepositoryMock) FindOne(userModel *models.User) (*models.User, error) {
	args := m.Called(userModel)
	return args.Get(0).(*models.User), args.Error(1)
}
func (m *userRepositoryMock) Create(userModel *models.User) error {
	args := m.Called()
	return args.Error(0)
}
func (m *userRepositoryMock) Update(userModel *models.User) error {
	args := m.Called()
	return args.Error(0)
}
func (m *userRepositoryMock) Delete(userModel *models.User) error {
	args := m.Called()
	return args.Error(0)
}
