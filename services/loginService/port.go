package loginService

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type Response struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
