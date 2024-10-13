package data_objects

type RegisterUserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
