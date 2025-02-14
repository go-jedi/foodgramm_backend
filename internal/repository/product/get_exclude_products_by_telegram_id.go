package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error) {
	r.logger.Debug("[GetExcludeProductsByTelegramID] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var uep product.UserExcludedProducts

	q := `
		SELECT *
		FROM user_excluded_products
		WHERE telegram_id = $1;
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, telegramID).Scan(
		&uep.ID, &uep.UserID, &uep.TelegramID,
		&uep.Products, &uep.CreatedAt, &uep.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get exclude products by telegram id", "err", err)
			return product.UserExcludedProducts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get exclude products by telegram id", "err", err)
		return product.UserExcludedProducts{}, fmt.Errorf("could not get exclude products by telegram id: %w", err)
	}

	return uep, nil
}
