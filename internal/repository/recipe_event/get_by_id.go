package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
)

func (r *repo) GetByID(ctx context.Context, recipeID int64) (recipeevent.Recipe, error) {
	r.logger.Debug("[get recipe event by id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var re recipeevent.Recipe

	q := `
		SELECT *
		FROM event_recipes
		WHERE id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		recipeID,
	).Scan(
		&re.ID, &re.TypeID, &re.Title,
		&re.Content, &re.CreatedAt, &re.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipe event by id", "err", err)
			return recipeevent.Recipe{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipe event by id", "err", err)
		return recipeevent.Recipe{}, fmt.Errorf("could not get recipe event by id: %w", err)
	}

	return re, nil
}
