package app

import "abcShop/services/registerService"

type SetupService struct {
	Register registerService.IRegisterService
}

func NewSetupService(repo *SetupRepository) *SetupService {
	register := registerService.NewRegisterService(repo.UserRepo)

	return &SetupService{
		Register: register,
	}
}
