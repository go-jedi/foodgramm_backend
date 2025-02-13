package payment

import "math/big"

//
// CREATE
//

type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Type       string `json:"type" validate:"required,oneof=rub stars"`
}

//
// GET LINK REQUEST
//

type GetLinkBody struct {
	TelegramID *big.Int `json:"telegram_id"`
	Type       string   `json:"type"`
}

type GetLinkResponse struct {
	URL string `json:"url"`
}
