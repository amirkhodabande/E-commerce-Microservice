package mocks

import (
	"github.com/ecommerce/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entities.UserEntity) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.UserEntity, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.UserEntity), args.Error(1)
}
