package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/utils"
)

func (s *serv) ExcludeProductsByID(ctx context.Context, dto product.ExcludeProductsByIDDTO) (product.ExcludeProductsByIDResponse, error) {
	dto.Products = utils.RemoveDuplicates(dto.Products)

	return s.productRepository.ExcludeProductsByID(ctx, dto)
}
