package types

type LoginUserSuccessResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	ID       string `json:"id"`
	Token    string `json:"token"`
}
