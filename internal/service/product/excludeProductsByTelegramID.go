package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/utils"
)

func (s *serv) ExcludeProductsByTelegramID(ctx context.Context, dto product.ExcludeProductsByTelegramIDDTO) (product.ExcludeProductsByTelegramIDResponse, error) {
	dto.Products = utils.RemoveDuplicates(dto.Products)

	return s.productRepository.ExcludeProductsByTelegramID(ctx, dto)
}
