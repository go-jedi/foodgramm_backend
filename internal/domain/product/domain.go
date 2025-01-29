package product

import "time"

type UserExcludedProducts struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	TelegramID string    `json:"telegram_id"`
	Products   []string  `json:"products"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
// ADD EXCLUDE PRODUCTS BY USER ID
//

type AddExcludeProductsByUserIDDTO struct {
	UserID   int64    `json:"user_id" validate:"required,gt=0"`
	Products []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

//
// ADD EXCLUDE PRODUCTS BY TELEGRAM ID
//

type AddExcludeProductsByTelegramIDDTO struct {
	TelegramID string   `json:"telegram_id" validate:"required,min=1"`
	Products   []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

//
// GET EXCLUDE PRODUCTS BY USER ID
//

type GetExcludeProductsByUserIDDTO struct {
	UserID int64 `json:"user_id" validate:"required,gt=0"`
}

//
// GET EXCLUDE PRODUCTS BY TELEGRAM ID
//

type GetExcludeProductsByTelegramIDDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
}
