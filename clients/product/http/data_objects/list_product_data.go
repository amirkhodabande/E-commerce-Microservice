package data_objects

import (
	"time"
)

type ListProductResponse struct {
	Status  int        `json:"status"`
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []*product `json:"data"`
}

type ListProductData struct {
	Name       string `query:"name,omitempty" url:"name,omitempty"`
	CategoryID uint   `query:"category_id,omitempty" url:"category_id,omitempty"`
	Page       int    `query:"page,omitempty" url:"page,omitempty"`
	Limit      int    `query:"limit,omitempty" url:"limit,omitempty"`
	SortBy     string `query:"sort_by,omitempty" url:"sort_by,omitempty"`
}

type product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (data *ListProductData) Validate() map[string]string {
	errors := make(map[string]string)

	if data.Page == 0 || data.Limit == 0 {
		errors["page"] = "Page and limit can't be 0"
	}

	if data.Limit > 100 {
		errors["limit"] = "Limit can't be greater than 100"
	}

	validSortOptions := map[string]bool{"most_expensive": true, "cheapest": true, "newest": true}
	if data.SortBy != "" && !validSortOptions[data.SortBy] {
		errors["sort_by"] = "Sort by can only be most_expensive, cheapest or newest"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
