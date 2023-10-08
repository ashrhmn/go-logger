package user

type AddUserDto struct {
	Username  string `json:"username" validate:"required,min=1,max=30"`
	Password  string `json:"password" validate:"required,min=1,max=40"`
	Email     string `json:"email" validate:"email,required,min=1,max=30"`
	FirstName string `json:"firstName" validate:"required,min=1,max=40"`
	LastName  string `json:"lastName" validate:"required,min=1,max=40"`
}
