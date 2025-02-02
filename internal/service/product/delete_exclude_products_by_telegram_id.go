package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (s *serv) DeleteExcludeProductsByTelegramID(ctx context.Context, telegramID string, prod string) (product.UserExcludedProducts, error) {
	return s.productRepository.DeleteExcludeProductsByTelegramID(ctx, telegramID, prod)
}
