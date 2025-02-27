package recipeevent

import "time"

type Recipe struct {
	ID        int64       `json:"id"`
	TypeID    int64       `json:"type_id"`
	Title     string      `json:"title"`
	Content   [][]Content `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// Content represents a content of recipe event in the systems.
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

//
// CREATE
//

// CreateDTO represents a creation recipe event dto in the system.
type CreateDTO struct {
	TypeID  int64       `json:"type_id" validate:"required,gt=0"`
	Title   string      `json:"title" validate:"required,min=1"`
	Content [][]Content `json:"content" validate:"min=1,dive,min=1,dive"`
}

//
// UPDATE
//

// UpdateDTO represents update recipe event dto in the system.
type UpdateDTO struct {
	ID     int64  `json:"id" validate:"required,gt=0"`
	TypeID int64  `json:"type_id" validate:"required,gt=0"`
	Title  string `json:"title" validate:"required,min=1"`
}
