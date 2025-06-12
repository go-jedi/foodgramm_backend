package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (r *repo) DeleteByID(ctx context.Context, id int64) (int64, error) {
	r.logger.Debug("[delete user by id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		DELETE FROM users
		WHERE id = $1;
	`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		id,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while delete user by id", "err", err)
			return 0, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to delete user by id", "err", err)
		return 0, fmt.Errorf("could not delete user by id: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return 0, apperrors.ErrNoRowsWereAffected
	}

	return id, nil
}
