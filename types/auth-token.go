package types

type AuthSession struct {
	Token       string   `json:"token"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
	ExpiresAt   int64    `json:"expiresAt"`
	CreatedAt   int64    `json:"createdAt"`
	Client      string   `json:"client"`
}
