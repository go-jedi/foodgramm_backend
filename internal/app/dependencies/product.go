package dependencies

import (
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	productRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/product"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	productService "github.com/go-jedi/foodgrammm-backend/internal/service/product"
)

func (d *Dependencies) ProductRepository() repository.ProductRepository {
	if d.productRepository == nil {
		d.productRepository = productRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.productRepository
}

func (d *Dependencies) ProductService() service.ProductService {
	if d.productService == nil {
		d.productService = productService.NewService(
			d.ProductRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.productService
}

func (d *Dependencies) ProductHandler() *product.Handler {
	if d.productHandler == nil {
		d.productHandler = product.NewHandler(
			d.ProductService(),
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.productHandler
}
