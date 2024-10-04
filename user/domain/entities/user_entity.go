package entities

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id        uint
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserEntity(id uint, email, password string) (*UserEntity, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password cannot be empty")
	}

	return &UserEntity{
		id:        id,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (entity *UserEntity) GetID() uint {
	return entity.id
}

func (entity *UserEntity) GetEmail() string {
	return entity.email
}

func (entity *UserEntity) GetPassword() string {
	return entity.password
}

func (entity *UserEntity) SetID(id uint) {
	entity.id = id
}

func (entity *UserEntity) HashPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(entity.password), 14)
	if err != nil {
		panic(err)
	}

	entity.password = string(bytes)
}

func (entity *UserEntity) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(entity.password), []byte(password))
	return err == nil
}
