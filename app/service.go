package app

import (
	"abcShop/services/loginService"
	"abcShop/services/registerService"
)

type SetupService struct {
	Register registerService.IRegisterService
	Login    loginService.ILoginService
}

func NewSetupService(repo *SetupRepository) *SetupService {
	register := registerService.NewRegisterService(repo.UserRepo)
	login := loginService.NewLoginService(repo.UserRepo)
	return &SetupService{
		Register: register,
		Login:    login,
	}
}
