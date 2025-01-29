package product

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	productRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/product"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	productService "github.com/go-jedi/foodgrammm-backend/internal/service/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Product struct {
	engine    *gin.Engine
	logger    *logger.Logger
	validator *validator.Validator
	db        *postgres.Postgres
	cache     *redis.Redis

	// product
	productRepository repository.ProductRepository
	productService    service.ProductService
	productHandler    *product.Handler
}

func NewProduct(
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
	db *postgres.Postgres,
	cache *redis.Redis,
) *Product {
	return &Product{
		engine:    engine,
		logger:    logger,
		validator: validator,
		db:        db,
		cache:     cache,
	}
}

func (p *Product) Init(ctx context.Context) {
	_ = p.ProductHandler(ctx)
}

func (p *Product) ProductRepository(_ context.Context) repository.ProductRepository {
	if p.productRepository == nil {
		p.productRepository = productRepository.NewRepository(
			p.logger,
			p.db,
		)
	}

	return p.productRepository
}

func (p *Product) ProductService(ctx context.Context) service.ProductService {
	if p.productService == nil {
		p.productService = productService.NewService(
			p.ProductRepository(ctx),
			p.logger,
			p.cache,
		)
	}

	return p.productService
}

func (p *Product) ProductHandler(ctx context.Context) *product.Handler {
	if p.productHandler == nil {
		p.productHandler = product.NewHandler(
			p.ProductService(ctx),
			p.engine,
			p.logger,
			p.validator,
		)
	}

	return p.productHandler
}
