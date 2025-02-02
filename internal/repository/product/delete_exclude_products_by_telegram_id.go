package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		UPDATE user_excluded_products SET
		    products = ARRAY_REMOVE(products, $1),
		    updated_at = NOW()
		WHERE telegram_id = $2
		RETURNING *;
	`

	var uep product.UserExcludedProducts

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		prod, telegramID,
	).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Products, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while delete exclude product by telegram id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to delete exclude product by telegram id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not delete exclude product by telegram id: %w", err)
	}

	return uep, nil
}
