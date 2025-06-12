package promocode

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	promoCodeRepository repository.PromoCodeRepository
	userRepository      repository.UserRepository
	logger              *logger.Logger
	cache               *redis.Redis
}

func NewService(
	promoCodeRepository repository.PromoCodeRepository,
	userRepository repository.UserRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.PromoCodeService {
	return &serv{
		promoCodeRepository: promoCodeRepository,
		userRepository:      userRepository,
		logger:              logger,
		cache:               cache,
	}
}
