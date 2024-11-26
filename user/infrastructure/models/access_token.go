package models

import "gorm.io/gorm"

type AccessToken struct {
	gorm.Model
	Name  string
	Token string
}
