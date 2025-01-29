package payment

//
// CREATE
//

type CreateDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
}

//
// CHECK STATUS
//

type CheckStatusDTO struct {
	TelegramID string `json:"telegram_id" validate:"required,min=1"`
}
