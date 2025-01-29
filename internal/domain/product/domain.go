package product

//
// ADD EXCLUDE PRODUCTS BY ID
//

type AddExcludeProductsByIDDTO struct {
	UserID   int64    `json:"user_id" validate:"required,gt=0"`
	Products []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

type AddExcludeProductsByIDResponse struct {
	UserID   int64    `json:"user_id" validate:"required,gt=0"`
	Products []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

//
// ADD EXCLUDE PRODUCTS BY TELEGRAM ID
//

type AddExcludeProductsByTelegramIDDTO struct {
	TelegramID string   `json:"telegram_id" validate:"required,min=1"`
	Products   []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}

type AddExcludeProductsByTelegramIDResponse struct {
	TelegramID string   `json:"telegram_id" validate:"required,min=1"`
	Products   []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
}
