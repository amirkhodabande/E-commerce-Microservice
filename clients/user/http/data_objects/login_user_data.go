package data_objects

type LoginUserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Status  int               `json:"status"`
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
	Token   string            `json:"token"`
}

func (data *LoginUserData) Validate() map[string]string {
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
