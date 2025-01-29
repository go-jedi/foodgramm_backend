package payment

//
// CREATE
//

type CreateDTO struct {
	TelegramID string `json:"telegram_id"`
}

//
// CHECK STATUS
//

type CheckStatusDTO struct {
	TelegramID string `json:"telegram_id"`
}
