package product

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/utils"
)

func (s *serv) AddExcludeProductsByID(ctx context.Context, dto product.AddExcludeProductsByIDDTO) (product.AddExcludeProductsByIDResponse, error) {
	dto.Products = utils.RemoveDuplicates(dto.Products)

	return s.productRepository.AddExcludeProductsByID(ctx, dto)
}
