package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (r *repo) All(ctx context.Context) ([]recipetypes.RecipeTypes, error) {
	r.logger.Debug("[get all recipe types] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM recipe_types
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipe types", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipe types", "err", err)
		return nil, fmt.Errorf("could not get recipe types: %w", err)
	}
	defer rows.Close()

	var recipeTypes []recipetypes.RecipeTypes

	for rows.Next() {
		var rc recipetypes.RecipeTypes

		if err := rows.Scan(
			&rc.ID, &rc.Title,
			&rc.CreatedAt, &rc.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all recipe types", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all recipe types: %w", err)
		}

		recipeTypes = append(recipeTypes, rc)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all recipe types", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all recipe types: %w", err)
	}

	return recipeTypes, nil
}
