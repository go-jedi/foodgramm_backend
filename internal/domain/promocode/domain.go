package promocode

import "time"

// PromoCode represents a promo code in the system.
type PromoCode struct {
	ID              int64      `json:"id"`
	Code            string     `json:"code"`
	DiscountPercent int        `json:"discount_percent"`
	MaxUsesAllowed  int        `json:"max_uses_allowed"`
	AmountUsed      int        `json:"amount_used"`
	IsReusable      bool       `json:"is_reusable"`
	ValidFrom       time.Time  `json:"valid_from"`
	ValidUntil      *time.Time `json:"valid_until"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

//
// CREATE
//

// CreateDTO represents the data required to create a new promo code.
type CreateDTO struct {
	Code            string     `json:"code" validate:"required,min=1"`
	DiscountPercent int        `json:"discount_percent" validate:"required,gt=0,lte=100"`
	MaxUsesAllowed  int        `json:"max_uses_allowed" validate:"required,gte=-1"`
	IsReusable      bool       `json:"is_reusable"`
	ValidUntil      *time.Time `json:"valid_until"`
}

//
// APPLY
//

type ApplyDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Code       string `json:"code" validate:"required,min=1"`
}

//
// CHECK
//

type CheckDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Code       string `json:"code" validate:"required,min=1"`
}
