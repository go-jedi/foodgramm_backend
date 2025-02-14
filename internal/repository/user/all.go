package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (r *repo) All(ctx context.Context) ([]user.User, error) {
	r.logger.Debug("[All] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM users
		ORDER BY id
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get users", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get users", "err", err)
		return nil, fmt.Errorf("could not get users: %w", err)
	}
	defer rows.Close()

	var usrs []user.User

	for rows.Next() {
		var u user.User

		if err := rows.Scan(
			&u.ID, &u.TelegramID, &u.Username,
			&u.FirstName, &u.LastName,
			&u.CreatedAt, &u.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all users", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all users: %w", err)
		}

		usrs = append(usrs, u)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all users", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return usrs, nil
}
