package data_objects

type RegisterUserParams struct {
	Username string `json:"user_name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
