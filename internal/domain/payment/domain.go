package payment

import "math/big"

type GetLinkBody struct {
	TelegramID *big.Int `json:"telegram_id"`
	Type       string   `json:"type"`
}

type GetLinkResponse struct {
	URL string `json:"url"`
}

type CheckStatusBody struct {
	TelegramID *big.Int `json:"telegram_id"`
}

type CheckStatusResponse struct {
	Value bool `json:"value"`
}

//
// CREATE
//

type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Type       string `json:"type" validate:"required,oneof=rub stars"`
}
