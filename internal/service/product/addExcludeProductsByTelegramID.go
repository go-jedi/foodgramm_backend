package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/utils"
)

func (s *serv) AddExcludeProductsByTelegramID(ctx context.Context, dto product.AddExcludeProductsByTelegramIDDTO) (product.UserExcludedProducts, error) {
	dto.Products = utils.RemoveDuplicates(dto.Products)

	return s.productRepository.AddExcludeProductsByTelegramID(ctx, dto)
}
