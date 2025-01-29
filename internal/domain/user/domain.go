package user

import (
	"time"

	jsoniter "github.com/json-iterator/go"
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
	Username   string `json:"username" validate:"required,min=1"`
	FirstName  string `json:"first_name" validate:"required,min=1"`
	LastName   string `json:"last_name" validate:"required,min=1"`
}

//
// LIST
//

type ListDTO struct {
	Page int `json:"page" validate:"required,gt=0"`
	Size int `json:"size" validate:"required,gt=0"`
}

type ListResponse struct {
	TotalCount  int                 `json:"total_count"`
	TotalPages  int                 `json:"total_pages"`
	CurrentPage int                 `json:"current_page"`
	Size        int                 `json:"size"`
	Data        jsoniter.RawMessage `json:"data"`
}

//
// EXISTS
//

type ExistsDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"required,min=1"`
}

//
// UPDATE
//

type UpdateDTO struct {
	ID         int64  `json:"id" validate:"required,gt=0"`
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Username   string `json:"username" validate:"required,min=1"`
	FirstName  string `json:"first_name" validate:"required,min=1"`
	LastName   string `json:"last_name" validate:"required,min=1"`
}

//
// DELETE
//

type DeleteDTO struct {
	ID         int64  `json:"id" validate:"required,gt=0"`
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
}

type DeleteResponse struct {
	ID         int64  `json:"id"`
	TelegramID string `json:"telegram_id"`
}
