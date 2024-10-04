package entities

import (
	"testing"

	"github.com/ecommerce/user/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewUserEntity(t *testing.T) {
	t.Run("should create a new user entity", func(t *testing.T) {
		user, err := entities.NewUserEntity(0, "test@example.com", "password")

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		user, err := entities.NewUserEntity(0, "", "password")

		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("should return error when password is empty", func(t *testing.T) {
		user, err := entities.NewUserEntity(0, "test@example.com", "")

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestHashPassword(t *testing.T) {
	user, err := entities.NewUserEntity(0, "test@example.com", "password")

	assert.NoError(t, err)
	assert.NotNil(t, user)

	user.HashPassword()

	t.Run("should hash the password", func(t *testing.T) {
		assert.NotEqual(t, "password", user.GetPassword())
	})

	t.Run("should compare the password", func(t *testing.T) {
		assert.True(t, user.ComparePassword("password"))
	})
}
