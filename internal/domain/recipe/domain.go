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

//
// GENERATE RECIPE
//

type GenerateRecipeDTO struct {
	TelegramID        string   `json:"telegram_id" validate:"required,min=1"`
	Type              int      `json:"type" validate:"required,gt=0,lte=5"`
	Allergies         string   `json:"allergies" validate:"required,min=1"`
	Products          []string `json:"products" validate:"required,min=1,max=50,dive,min=1"`
	Name              *string  `json:"name" validate:"omitempty,min=1"`
	AmountCalories    *int     `json:"amount_calories" validate:"omitempty,gt=0"`
	AvailableProducts []string `json:"available_products" validate:"omitempty,min=1,max=50,dive,min=1"`
}
