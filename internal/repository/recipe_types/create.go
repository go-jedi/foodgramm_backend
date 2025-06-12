package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (r *repo) Create(ctx context.Context, dto recipetypes.CreateDTO) (recipetypes.RecipeTypes, error) {
	r.logger.Debug("[create a new recipe type] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO recipe_types(
			title
		) VALUES($1)
		RETURNING *;
	`

	var nrt recipetypes.RecipeTypes

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Title,
	).Scan(
		&nrt.ID, &nrt.Title,
		&nrt.CreatedAt, &nrt.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the recipe type", "err", err)
			return recipetypes.RecipeTypes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create the recipe type", "err", err)
		return recipetypes.RecipeTypes{}, fmt.Errorf("could not create recipe type: %w", err)
	}

	return nrt, nil
}
