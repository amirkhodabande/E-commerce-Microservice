package persistence

import (
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/infrastructure/models"
	"gorm.io/gorm"
)

type SqlAccessTokenRepository struct {
	db *gorm.DB
}

func NewSqlAccessTokenRepository(db *gorm.DB) *SqlAccessTokenRepository {
	return &SqlAccessTokenRepository{
		db: db,
	}
}

func (r *SqlAccessTokenRepository) Create(accessToken *entities.AccessTokenEntity) error {
	modelAccessToken := &models.AccessToken{
		Name:  accessToken.GetName(),
		Token: accessToken.GetToken(),
	}

	result := r.db.Create(modelAccessToken)

	accessToken.SetID(modelAccessToken.ID)

	return result.Error
}
