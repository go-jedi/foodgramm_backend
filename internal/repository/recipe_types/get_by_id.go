package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (r *repo) GetByID(ctx context.Context, recipeTypeID int64) (recipetypes.RecipeTypes, error) {
	r.logger.Debug("[get recipe type by id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var rt recipetypes.RecipeTypes

	q := `
		SELECT *
		FROM recipe_types
		WHERE id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q, recipeTypeID,
	).Scan(
		&rt.ID, &rt.Title,
		&rt.CreatedAt, &rt.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipe type by id", "err", err)
			return recipetypes.RecipeTypes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipe type by id", "err", err)
		return recipetypes.RecipeTypes{}, fmt.Errorf("could not get recipe type by id: %w", err)
	}

	return rt, nil
}
