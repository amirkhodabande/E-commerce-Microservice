package contracts

import "github.com/ecommerce/user/domain/entities"

type UserRepository interface {
	Create(user *entities.UserEntity) error
}
