package userblacklist

import "time"

// UsersBlackList represents a users blacklist entry in the system.
type UsersBlackList struct {
	ID                 int64     `json:"id"`
	TelegramID         string    `json:"telegram_id"`
	BanTimestamp       time.Time `json:"ban_timestamp"`
	BanReason          string    `json:"ban_reason"`
	BannedByTelegramID string    `json:"banned_by_telegram_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

//
// BLOCK USER DTO
//

// BlockUserDTO represents a block user dto in the system.
type BlockUserDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	BanReason  string `json:"ban_reason" validate:"required,min=1"`
}
