package product

import "time"

// UserExcludedProducts represents the response structure for user excluded products.
type UserExcludedProducts struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	TelegramID string    `json:"telegram_id"`
	Allergies  string    `json:"allergies"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
// ADD ALLERGIES BY TELEGRAM ID
//

// AddAllergiesByTelegramIDDTO represents the request body for adding allergies.
type AddAllergiesByTelegramIDDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Allergies  string `json:"allergies" validate:"required,min=1"`
}
