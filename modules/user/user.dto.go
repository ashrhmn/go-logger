package user

type AddUserDto struct {
	Username  string `json:"username" validate:"required,min=1,max=30"`
	Password  string `json:"password" validate:"required,min=1,max=40"`
	Email     string `json:"email" validate:"email,required,min=1,max=30"`
	FirstName string `json:"firstName" validate:"required,min=1,max=40"`
	LastName  string `json:"lastName" validate:"required,min=1,max=40"`
}

type UpdateUserDto struct {
	Username    string   `json:"username,omitempty" validate:"omitempty,required,min=1,max=30"`
	Password    string   `json:"password,omitempty" validate:"omitempty,required,min=1,max=40"`
	Email       string   `json:"email,omitempty" validate:"omitempty,required,email,min=1,max=30"`
	FirstName   string   `json:"firstName,omitempty" validate:"omitempty,required,min=1,max=40"`
	LastName    string   `json:"lastName,omitempty" validate:"omitempty,required,min=1,max=40"`
	Permissions []string `json:"permissions"`
}
