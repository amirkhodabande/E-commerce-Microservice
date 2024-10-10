package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/infrastructure/models"
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	app, repoContainer := setup(t)

	t.Run("should register user", func(t *testing.T) {
		registerUserParams := data_objects.RegisterUserParams{
			Email:    "test@example.com",
			Password: "password",
		}
		b, _ := json.Marshal(registerUserParams)

		request := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(b))
		request.Header.Add("Content-Type", "application-json")

		response, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, http.StatusCreated, response.StatusCode)

		var user models.User
		result := repoContainer.DB.Where("email = ?", registerUserParams.Email).First(&user)
		if result.Error != nil {
			t.Error(result.Error)
		}
		assert.Equal(t, registerUserParams.Email, user.Email)
	})

	t.Run("should return error when email is already taken", func(t *testing.T) {
		registerUserParams := data_objects.RegisterUserParams{
			Email:    "test@example.com",
			Password: "password",
		}
		b, _ := json.Marshal(registerUserParams)

		user, err := entities.NewUserEntity(0, registerUserParams.Email, registerUserParams.Password)
		if err != nil {
			t.Error(err)
		}

		modelUser := &models.User{
			Email:    user.GetEmail(),
			Password: user.GetPassword(),
		}

		result := repoContainer.DB.Create(modelUser)
		if result.Error != nil {
			t.Error(result.Error)
		}

		request := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(b))
		request.Header.Add("Content-Type", "application-json")

		response, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})
}

func TestLogin(t *testing.T) {
	app, repoContainer := setup(t)
	// create user
	user, err := entities.NewUserEntity(0, "test@example.com", "password")
	if err != nil {
		t.Error(err)
	}
	user.HashPassword()
	repoContainer.DB.Create(&models.User{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	})

	t.Run("should login user", func(t *testing.T) {
		loginUserParams := data_objects.LoginUserParams{
			Email:    "test@example.com",
			Password: "password",
		}
		b, _ := json.Marshal(loginUserParams)

		request := httptest.NewRequest("POST", "/api/user/login", bytes.NewReader(b))
		request.Header.Add("Content-Type", "application-json")

		response, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}
