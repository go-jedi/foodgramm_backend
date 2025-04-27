package admin

import "time"

// Admin represents a admin in the system.
type Admin struct {
	ID         int64     `json:"id"`
	TelegramID string    `json:"telegram_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
// ADD ADMIN USER DTO
//

// AddAdminUserDTO represents the data required to add admin user.
type AddAdminUserDTO struct {
	TelegramID string `json:"telegram_id"`
}
