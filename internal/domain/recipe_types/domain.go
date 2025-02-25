package recipetypes

import "time"

// RecipeTypes represents a recipe types entry in the system.
type RecipeTypes struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//
// CREATE
//

// CreateDTO represents a creation recipe type dto in the system.
type CreateDTO struct {
	Title string `json:"title" validate:"required,min=1"`
}

//
// UPDATE
//

// UpdateDTO represents update recipe type dto in the system.
type UpdateDTO struct {
	ID    int64  `json:"id" validate:"required,gt=0"`
	Title string `json:"title" validate:"required,min=1"`
}
