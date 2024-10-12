package application

import (
	"net/http"

	"github.com/ecommerce/user/domain/contracts"
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/domain/services"
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
)

type AuthService struct {
	contracts.UserRepository
}

func NewAuthService(repository contracts.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: repository,
	}
}

func (service *AuthService) Register(user *entities.UserEntity) (string, error) {
	if !services.IsEmailUnique(service.UserRepository, user.GetEmail()) {
		return "", data_objects.NewApiError(http.StatusBadRequest, "user already exists")
	}

	user.HashPassword()

	service.UserRepository.Create(user)

	token, err := services.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) Login(userEntity *entities.UserEntity) (string, error) {
	user, err := service.UserRepository.FindByEmail(userEntity.GetEmail())
	if err != nil {
		return "", err
	}

	if !user.ComparePassword(userEntity.GetPassword()) {
		return "", data_objects.NewApiError(http.StatusUnauthorized, "invalid password")
	}

	token, err := services.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
