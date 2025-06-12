package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

func (r *repo) AllByTypeID(ctx context.Context, typeID int64) ([]recipeevent.Recipe, error) {
	r.logger.Debug("[get all recipes event by type id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM event_recipes
		WHERE type_id = $1
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q, typeID)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipes event by type id", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipes event by type id", "err", err)
		return nil, fmt.Errorf("could not get recipes event by type id: %w", err)
	}
	defer rows.Close()

	var recipesEvent []recipeevent.Recipe

	for rows.Next() {
		var re recipeevent.Recipe

		if err := rows.Scan(
			&re.ID, &re.TypeID, &re.Title,
			&re.Content, &re.CreatedAt, &re.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all recipes event by type id", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all recipes event by type id: %w", err)
		}

		recipesEvent = append(recipesEvent, re)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all recipes event by type id", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all recipes event by type id: %w", err)
	}

	return recipesEvent, nil
}
