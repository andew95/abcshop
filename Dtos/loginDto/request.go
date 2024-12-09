package loginDto

import "abcShop/services/loginService"

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Request) ToLoginServiceRequest() loginService.Request {
	return loginService.Request{
		Email:    r.Email,
		Password: r.Password,
	}
}
