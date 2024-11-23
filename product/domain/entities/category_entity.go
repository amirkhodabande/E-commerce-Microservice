package entities

import (
	"time"
)

type CategoryEntity struct {
	Id          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Children    []CategoryEntity `json:"children"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   time.Time        `json:"deleted_at"`
}
