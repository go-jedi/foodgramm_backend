package user

import (
	"encoding/json"
	"time"
)

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

type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
}

//
// LIST
//

type ListDTO struct {
	Page int `json:"page" validate:"required,min=1"`
	Size int `json:"size" validate:"required,min=1"`
}

type ListResponse struct {
	TotalCount  int             `json:"total_count"`
	TotalPages  int             `json:"total_pages"`
	CurrentPage int             `json:"current_page"`
	Size        int             `json:"size"`
	Data        json.RawMessage `json:"data"`
}

//
// EXISTS
//

type ExistsDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"required"`
}

//
// UPDATE
//

type UpdateDTO struct {
	ID         int64  `json:"id"`
	TelegramID string `json:"telegram_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}
