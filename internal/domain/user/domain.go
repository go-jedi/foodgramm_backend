package user

import "time"

type User struct {
	ID         int64
	TelegramID int64
	Username   string
	FirstName  string
	LastName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//
// CREATE
//

type CreateDTO struct {
	TelegramID int64  `json:"telegram_id" validate:"required,gt=0"`
	Username   string `json:"username" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
}
