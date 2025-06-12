package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (r *repo) DeleteByTelegramID(ctx context.Context, telegramID string) (string, error) {
	r.logger.Debug("[delete user by telegram id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		DELETE FROM users
		WHERE telegram_id = $1;
	`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		telegramID,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while delete user by telegramID", "err", err)
			return "", fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to delete user by telegramID", "err", err)
		return "", fmt.Errorf("could not delete user by telegramID: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return "", apperrors.ErrNoRowsWereAffected
	}

	return telegramID, nil
}
