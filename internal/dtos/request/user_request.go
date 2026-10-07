package request

type UserRequest struct {
	ID       string
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequestList struct {
	Page   int
	Limit  int
	Search string
}
