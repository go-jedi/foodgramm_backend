package user

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) GetUserCount(ctx context.Context) (int64, error) {
	r.logger.Debug("[get user count] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	cnt := int64(0)

	q := `SELECT COUNT(*) FROM users;`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
	).Scan(&cnt); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get user count", "err", err)
			return 0, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get user count", "err", err)
		return 0, fmt.Errorf("could not get user count: %w", err)
	}

	return cnt, nil
}
