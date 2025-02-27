package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
)

func (r *repo) All(ctx context.Context) ([]recipeevent.Recipe, error) {
	r.logger.Debug("[get all recipes event] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM event_recipes
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipes event", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipes event", "err", err)
		return nil, fmt.Errorf("could not get recipes event: %w", err)
	}
	defer rows.Close()

	var recipesEvent []recipeevent.Recipe

	for rows.Next() {
		var re recipeevent.Recipe

		if err := rows.Scan(
			&re.ID, &re.TypeID, &re.Title,
			&re.Content, &re.CreatedAt, &re.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all recipes event", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all recipes event: %w", err)
		}

		recipesEvent = append(recipesEvent, re)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all recipes event", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all recipes event: %w", err)
	}

	return recipesEvent, nil
}
