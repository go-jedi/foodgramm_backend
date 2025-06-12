package userblacklist

import (
	"context"
	"errors"
	"fmt"
	"time"

	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
)

func (r *repo) BlockUser(ctx context.Context, dto userblacklist.BlockUserDTO, bannedByTelegramID string) (userblacklist.UsersBlackList, error) {
	r.logger.Debug("[block user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO users_blacklist(
		    telegram_id,
		    ban_reason,
		    banned_by_telegram_id
		) VALUES($1, $2, $3)
		RETURNING *;
	`

	var nubl userblacklist.UsersBlackList

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.TelegramID, dto.BanReason, bannedByTelegramID,
	).Scan(
		&nubl.ID, &nubl.TelegramID, &nubl.BanTimestamp,
		&nubl.BanReason, &nubl.BannedByTelegramID,
		&nubl.CreatedAt, &nubl.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while block user", "err", err)
			return userblacklist.UsersBlackList{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to block user", "err", err)
		return userblacklist.UsersBlackList{}, fmt.Errorf("could not block user: %w", err)
	}

	return nubl, nil
}
