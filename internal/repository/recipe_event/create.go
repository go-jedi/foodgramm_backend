package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

func (r *repo) Create(ctx context.Context, dto recipeevent.CreateDTO) (recipeevent.Recipe, error) {
	r.logger.Debug("[create a new recipe event] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO event_recipes(
		    type_id,
		    title,
		    content
		) VALUES($1, $2, $3)
		RETURNING *;
	`

	var nre recipeevent.Recipe

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.TypeID, dto.Title, dto.Content,
	).Scan(
		&nre.ID, &nre.TypeID, &nre.Title,
		&nre.Content, &nre.CreatedAt, &nre.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the recipe event", "err", err)
			return recipeevent.Recipe{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create the recipe event", "err", err)
		return recipeevent.Recipe{}, fmt.Errorf("could not create recipe event: %w", err)
	}

	return nre, nil
}
