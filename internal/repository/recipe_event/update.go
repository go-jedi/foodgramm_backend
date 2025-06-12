package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

func (r *repo) Update(ctx context.Context, dto recipeevent.UpdateDTO) (recipeevent.Recipe, error) {
	r.logger.Debug("[update recipe event] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var re recipeevent.Recipe

	q := `
		UPDATE event_recipes SET
		    type_id = $1,
			title = $2,
			updated_at = NOW()
		WHERE id = $3
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.TypeID, dto.Title, dto.ID,
	).Scan(
		&re.ID, &re.TypeID, &re.Title,
		&re.Content, &re.CreatedAt, &re.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while update recipe event", "err", err)
			return recipeevent.Recipe{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to update recipe event", "err", err)
		return recipeevent.Recipe{}, fmt.Errorf("could not update recipe event: %w", err)
	}

	return re, nil
}
