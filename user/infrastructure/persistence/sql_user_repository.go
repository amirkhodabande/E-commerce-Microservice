package persistence

import (
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/infrastructure/models"
	"gorm.io/gorm"
)

type SqlUserRepository struct {
	db *gorm.DB
}

func NewSqlUserRepository(db *gorm.DB) *SqlUserRepository {
	return &SqlUserRepository{
		db: db,
	}
}

func (r *SqlUserRepository) FindByEmail(email string) (*entities.UserEntity, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	userEntity, err := entities.NewUserEntity(user.ID, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (r *SqlUserRepository) Create(user *entities.UserEntity) error {
	modelUser := &models.User{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}

	result := r.db.Create(modelUser)

	user.SetID(modelUser.ID)

	return result.Error
}
