package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) AddAllergiesByTelegramID(ctx context.Context, dto product.AddAllergiesByTelegramIDDTO) (product.UserExcludedProducts, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var uep product.UserExcludedProducts

	q := `
		UPDATE user_excluded_products SET
			allergies = $1,
			updated_at = NOW()
		WHERE telegram_id = $2
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Allergies, dto.TelegramID,
	).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Allergies, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add allergies by telegram id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add allergies by telegram id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not add allergies by telegram id: %w", err)
	}

	return uep, nil
}
