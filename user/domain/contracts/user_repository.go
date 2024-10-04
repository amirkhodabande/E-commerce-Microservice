package contracts

import "github.com/ecommerce/user/domain/entities"

type UserRepository interface {
	Create(user *entities.UserEntity) error
	FindByEmail(email string) (*entities.UserEntity, error)
}
