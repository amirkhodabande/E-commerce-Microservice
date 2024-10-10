package data_objects

type LoginUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
