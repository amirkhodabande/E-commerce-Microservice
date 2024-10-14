package data_objects

type RegisterUserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Status  int               `json:"status"`
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Token   string            `json:"token"`
	Errors  map[string]string `json:"errors"`
}

func (data *RegisterUserData) Validate() map[string]string {
	errors := make(map[string]string)

	if data.Email == "" {
		errors["email"] = "Email is required"
	}

	if data.Password == "" {
		errors["password"] = "Password is required"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
