package user

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// User represents a user in the system.
type User struct {
	ID         int64     `json:"id"`
	TelegramID string    `json:"telegram_id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
// CREATE
//

// CreateDTO represents the data required to create a new user.
type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"omitempty,min=1"`
	FirstName  string `json:"first_name" validate:"required,min=1"`
	LastName   string `json:"last_name" validate:"omitempty,min=1"`
}

//
// LIST
//

// ListDTO represents the data required for pagination.
type ListDTO struct {
	Page int `json:"page" validate:"required,gt=0"`
	Size int `json:"size" validate:"required,gt=0"`
}

// ListResponse represents the response for a list of users with pagination.
type ListResponse struct {
	TotalCount  int                 `json:"total_count"`
	TotalPages  int                 `json:"total_pages"`
	CurrentPage int                 `json:"current_page"`
	Size        int                 `json:"size"`
	Data        jsoniter.RawMessage `json:"data"`
}

// ListResponseSwagger represents the response for a list of users with pagination for swagger.
type ListResponseSwagger struct {
	TotalCount  int    `json:"total_count"`
	TotalPages  int    `json:"total_pages"`
	CurrentPage int    `json:"current_page"`
	Size        int    `json:"size"`
	Data        []User `json:"data"`
}

//
// EXISTS
//

// ExistsDTO represents the data required to check if a user exists.
type ExistsDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"omitempty,min=1"`
}

//
// UPDATE
//

// UpdateDTO represents the data required to update a user.
type UpdateDTO struct {
	ID         int64  `json:"id" validate:"required,gt=0"`
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"omitempty,min=1"`
	FirstName  string `json:"first_name" validate:"required,min=1"`
	LastName   string `json:"last_name" validate:"omitempty,min=1"`
}
