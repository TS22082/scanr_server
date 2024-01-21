package types

type UserSuccessResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}
