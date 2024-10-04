package services

import (
	"os"
	"time"

	"github.com/ecommerce/user/domain/entities"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *entities.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.GetID(),
		"email":   user.GetEmail(),
		"expires": time.Now().Add(time.Hour * 4),
	})
	secret := os.Getenv("JWT_SECRET")
	tokenStr, _ := token.SignedString([]byte(secret))

	return tokenStr, nil
}
