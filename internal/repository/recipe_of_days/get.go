package recipeofdays

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
)

func (r *repo) Get(ctx context.Context) (recipeofdays.Recipe, error) {
	r.logger.Debug("[get recipe of the day] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var rod recipeofdays.Recipe

	q := `
		SELECT *
		FROM recipe_of_days
		ORDER BY id DESC
		LIMIT 1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
	).Scan(
		&rod.ID, &rod.Title, &rod.Lifehack,
		&rod.Content, &rod.CreatedAt, &rod.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipe of days", "err", err)
			return recipeofdays.Recipe{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipe of days", "err", err)
		return recipeofdays.Recipe{}, fmt.Errorf("could not get recipe of days: %w", err)
	}

	return rod, nil
}
