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

// InCache represents a recipe entry in the cache.
type InCache struct {
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
	Calories          string   `json:"calories"`
	Bzhu              string   `json:"bzhu"`
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

//
// GENERATE RECIPE DTO
//

// GenerateRecipeDTO represents a generate recipe dto in the system.
type GenerateRecipeDTO struct {
	TelegramID            string   `json:"telegram_id" validate:"required,min=1"`
	Type                  int      `json:"type" validate:"required,gt=0,lte=4"`
	Products              []string `json:"products" validate:"required,max=50,dive,min=1"`
	NonConsumableProducts *string  `json:"non_consumable_products" validate:"omitempty,min=0"`
	Name                  *string  `json:"name" validate:"omitempty,min=1"`
	AmountCalories        *int     `json:"amount_calories" validate:"omitempty,gt=0"`
	AvailableProducts     []string `json:"available_products" validate:"omitempty,min=1,max=50,dive,min=1"`
}

//
// GENERATE RECIPE RESPONSE
//

// GenerateRecipeResponse represents a generate recipe response in the system.
type GenerateRecipeResponse struct {
	TelegramID string      `json:"telegram_id"`
	Title      string      `json:"title"`
	Content    [][]Content `json:"content"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
