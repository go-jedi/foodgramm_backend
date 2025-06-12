package userblacklist

import (
	"context"
	"errors"
	"fmt"
	"time"

	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
)

func (r *repo) AllBannedUsers(ctx context.Context) ([]userblacklist.UsersBlackList, error) {
	r.logger.Debug("[get all banned users] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM users_blacklist
		ORDER BY id;
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get banned users", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get banned users", "err", err)
		return nil, fmt.Errorf("could not get banned users: %w", err)
	}
	defer rows.Close()

	var usbl []userblacklist.UsersBlackList

	for rows.Next() {
		var ubl userblacklist.UsersBlackList

		if err := rows.Scan(
			&ubl.ID, &ubl.TelegramID, &ubl.BanTimestamp,
			&ubl.BanReason, &ubl.BannedByTelegramID,
			&ubl.CreatedAt, &ubl.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get banned users", "err", err)
			return nil, fmt.Errorf("failed to scan row to get banned users: %w", err)
		}

		usbl = append(usbl, ubl)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get banned users", "err", rows.Err())
		return nil, fmt.Errorf("failed to get banned users: %w", err)
	}

	return usbl, nil
}
