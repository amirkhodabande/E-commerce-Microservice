package persistence

import (
	"fmt"

	"github.com/ecommerce/user/domain/entities"
)

type SqlUserRepository struct {
	// db *gorm.DB
}

func NewSqlUserRepository() *SqlUserRepository {
	return &SqlUserRepository{}
}

func (r *SqlUserRepository) Create(user *entities.UserEntity) error {
	fmt.Printf("%+v \n", user)
	return nil
}
