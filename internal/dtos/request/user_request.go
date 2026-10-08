package request

type UserRequest struct {
	ID       string
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserListRequest struct {
	Page   int
	Limit  int
	Search string
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Name                 string `json:"name" binding:"required"`
	Username             string `json:"username" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
