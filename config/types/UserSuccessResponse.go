package types

type UserSuccessResponse struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Verified bool   `json:"verified"`
}
