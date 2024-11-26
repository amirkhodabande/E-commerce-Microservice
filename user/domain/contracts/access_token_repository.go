package contracts

import "github.com/ecommerce/user/domain/entities"

type AccessTokenRepository interface {
	Create(accessToken *entities.AccessTokenEntity) error
}
