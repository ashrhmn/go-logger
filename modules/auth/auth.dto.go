package auth

type LoginInput struct {
	UsernameOrEmail string `json:"usernameOrEmail" validate:"required,min=1,max=30"`
	Password        string `json:"password" validate:"required,min=1,max=80"`
}

type SignUpInput struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Email    string `json:"email" validate:"email,required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=6,max=40"`
	Name     string `json:"name" validate:"min=1,max=30"`
}
