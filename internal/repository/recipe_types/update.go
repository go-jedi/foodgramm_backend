package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (r *repo) Update(ctx context.Context, dto recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error) {
	r.logger.Debug("[update recipe type] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var rt recipetypes.RecipeTypes

	q := `
		UPDATE recipe_types SET
			title = $1,
			updated_at = NOW()
		WHERE id = $2
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Title, dto.ID,
	).Scan(
		&rt.ID, &rt.Title,
		&rt.CreatedAt, &rt.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while update recipe type", "err", err)
			return recipetypes.RecipeTypes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to update recipe type", "err", err)
		return recipetypes.RecipeTypes{}, fmt.Errorf("could not update recipe type: %w", err)
	}

	return rt, nil
}
