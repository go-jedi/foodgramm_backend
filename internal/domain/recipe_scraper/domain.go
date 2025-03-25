package recipescraper

type GetBody struct {
	TelegramID            string   `json:"telegram_id"`
	Type                  int      `json:"type"`
	NonConsumableProducts []string `json:"non_consumable_products"`
}
