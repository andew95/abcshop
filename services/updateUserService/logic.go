package updateUserService

import (
	"abcShop/models"
	"time"
)

func (r *Request) ToModel(userModel *models.User) *models.User {
	userModel.Name = r.Name
	userModel.Email = r.Email
	userModel.Address = r.Address
	return userModel
}

func ToResponse(userModel *models.User) Response {
	return Response{
		User: UserResponse{
			Id:        userModel.Id.String(),
			Username:  userModel.Username,
			Name:      userModel.Name,
			Email:     userModel.Email,
			Address:   userModel.Address,
			Status:    userModel.Status.ToString(),
			CreatedAt: userModel.CreatedAt.Format(time.ANSIC),
		},
	}
}
