package services

import "github.com/ecommerce/user/domain/contracts"

func IsEmailUnique(repository contracts.UserRepository, email string) bool {
	user, _ := repository.FindByEmail(email)

	return user == nil
}
