package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

func (s *serv) GetAllergiesByTelegramID(ctx context.Context, telegramID string) (product.UserExcludedProducts, error) {
	return s.productRepository.GetAllergiesByTelegramID(ctx, telegramID)
}
