package product

//
// EXCLUDE PRODUCTS BY ID
//

type ExcludeProductsByIDDTO struct {
	UserID   int64    `json:"user_id" validate:"required,gt=0"`
	Products []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

type ExcludeProductsByIDResponse struct {
	UserID   int64    `json:"user_id" validate:"required,gt=0"`
	Products []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

//
// EXCLUDE PRODUCTS BY TELEGRAM ID
//

type ExcludeProductsByTelegramIDDTO struct {
	TelegramID string   `json:"telegram_id" validate:"required,min=1"`
	Products   []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

type ExcludeProductsByTelegramIDResponse struct {
	TelegramID string   `json:"telegram_id" validate:"required,min=1"`
	Products   []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}
