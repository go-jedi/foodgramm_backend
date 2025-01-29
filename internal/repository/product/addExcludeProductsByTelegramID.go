package product

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (r *repo) AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.AddExcludeProductsByTelegramIDResponse, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		UPDATE user_excluded_products_table SET
			products = ARRAY_CAT(products, (
            	SELECT ARRAY_AGG(DISTINCT product)
            	FROM UNNEST($1::TEXT[]) AS product
            	WHERE product NOT IN (SELECT UNNEST(products))
        	)
		) 
        WHERE telegram_id = $2
		RETURNING telegram_id, products;
	`

	var ep product.AddExcludeProductsByTelegramIDResponse

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Products, dto.TelegramID,
	).Scan(
		&ep.TelegramID, &ep.Products,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add exclude products by telegram id", "err", err)
			return product.AddExcludeProductsByTelegramIDResponse{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add exclude products by id", "err", err)
		return product.AddExcludeProductsByTelegramIDResponse{}, fmt.Errorf("could not add exclude products by telegram id: %w", err)
	}

	return ep, nil
}
