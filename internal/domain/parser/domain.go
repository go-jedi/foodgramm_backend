package parser

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

type ParsedRecipe struct {
	TelegramID string      `json:"telegram_id"`
	Title      string      `json:"title"`
	Content    [][]Content `json:"content"`
}
