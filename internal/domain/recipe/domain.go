package recipe

import (
	"time"
)

// Recipes represents a recipes entry in the system.
type Recipes struct {
	ID         int64       `json:"id"`
	TelegramID string      `json:"telegram_id"`
	Title      string      `json:"title"`
	Content    [][]Content `json:"content"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// Content represents a content of recipes in the systems.
type Content struct {
	ID                int64    `json:"id"`
	Type              string   `json:"type"`
	Title             string   `json:"title"`
	RecipePreparation string   `json:"recipe_preparation"`
	Ingredients       []string `json:"ingredients"`
	MethodPreparation []string `json:"method_preparation"`
}

// UserFreeRecipes represents a user free recipes in the system.
type UserFreeRecipes struct {
	ID                 int64     `json:"id"`
	TelegramID         string    `json:"telegram_id"`
	FreeRecipesAllowed int       `json:"free_recipes_allowed"`
	FreeRecipesUsed    int       `json:"free_recipes_used"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
