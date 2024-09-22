package application

import (
	"reflect"

	"github.com/ecommerce/user/domain/contracts"
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/infrastructure/persistence"
)

type AuthService struct {
	contracts.UserRepository
}

func NewAuthService() *AuthService {
	container := persistence.NewContainer()

	userRepository, err := container.Resolve(reflect.TypeOf((*contracts.UserRepository)(nil)))
	if err != nil {
		panic(err)
	}

	return &AuthService{
		UserRepository: userRepository.(contracts.UserRepository),
	}
}

func (service *AuthService) Register(registerData *entities.UserEntity) error {
	service.UserRepository.Create(registerData)

	return nil
}
