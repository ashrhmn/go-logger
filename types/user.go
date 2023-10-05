package types

type User struct {
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Permissions []string `json:"permissions"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   int64    `json:"updatedAt"`
	CreatedBy   string   `json:"createdBy"`
	UpdatedBy   string   `json:"updatedBy"`
	DeletedAt   int64    `json:"deletedAt"`
	DeletedBy   string   `json:"deletedBy"`
}
