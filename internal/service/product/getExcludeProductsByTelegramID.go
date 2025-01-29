package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (s *serv) GetExcludeProductsByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error) {
	return s.productRepository.GetExcludeProductsByTelegramID(ctx, telegramID)
}
