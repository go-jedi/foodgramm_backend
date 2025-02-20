package payment

import "math/big"

//
// CREATE
//

type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
	Type       string `json:"type" validate:"required,oneof=rub stars"`
	Price      int64  `json:"price" validate:"required,gt=0"`
}

//
// GET LINK REQUEST
//

type GetLinkBody struct {
	TelegramID *big.Int `json:"telegram_id"`
	Type       string   `json:"type"`
	Price      int64    `json:"price"`
}

type GetLinkResponse struct {
	URL string `json:"url"`
}
