package recipe

import "time"

// UserFreeRecipes represents a user free recipes in the system.
type UserFreeRecipes struct {
	ID                 int64     `json:"id"`
	TelegramID         string    `json:"telegram_id"`
	FreeRecipesAllowed int       `json:"free_recipes_allowed"`
	FreeRecipesUsed    int       `json:"free_recipes_used"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
