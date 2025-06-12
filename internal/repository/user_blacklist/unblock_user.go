package userblacklist

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (r *repo) UnblockUser(ctx context.Context, telegramID string) (string, error) {
	r.logger.Debug("[unblock user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		DELETE FROM users_blacklist
		WHERE telegram_id = $1;
	`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		telegramID,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while unblock user", "err", err)
			return "", fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to unblock user", "err", err)
		return "", fmt.Errorf("could not unblock user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return "", apperrors.ErrNoRowsWereAffected
	}

	return telegramID, nil
}
