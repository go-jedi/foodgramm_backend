package recipeofdays

import "time"

// Recipe represents a recipe of days entry in the system.
type Recipe struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Content     [][]Content `json:"content"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// Content represents a content of recipe of days in the systems.
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
