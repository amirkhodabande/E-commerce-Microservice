package entities

import (
	"errors"
	"time"
)

type AccessTokenEntity struct {
	id        uint
	name      string
	token     string
	createdAt time.Time
	updatedAt time.Time
}

func NewAccessTokenEntity(id uint, name string, token string) (*AccessTokenEntity, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if token == "" {
		return nil, errors.New("token cannot be empty")
	}

	return &AccessTokenEntity{
		id:        id,
		name:      name,
		token:     token,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (entity *AccessTokenEntity) GetID() uint {
	return entity.id
}

func (entity *AccessTokenEntity) GetName() string {
	return entity.name
}

func (entity *AccessTokenEntity) GetToken() string {
	return entity.token
}

func (entity *AccessTokenEntity) SetID(id uint) {
	entity.id = id
}
