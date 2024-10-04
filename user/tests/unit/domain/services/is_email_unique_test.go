package services

import (
	"testing"

	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/domain/services"
	"github.com/ecommerce/user/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestIsEmailUnique(t *testing.T) {
	t.Run("should return true if email is unique", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("FindByEmail", "test@example.com").Return((*entities.UserEntity)(nil), nil)

		unique := services.IsEmailUnique(mockRepo, "test@example.com")
		assert.True(t, unique)
	})

	t.Run("should return false if email is not unique", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("FindByEmail", "test@example.com").Return(&entities.UserEntity{}, nil)

		unique := services.IsEmailUnique(mockRepo, "test@example.com")
		assert.False(t, unique)
	})
}
