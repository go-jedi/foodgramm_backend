package subscription

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	subscriptionRepository repository.SubscriptionRepository
	userRepository         repository.UserRepository
	logger                 *logger.Logger
	cache                  *redis.Redis
}

func NewService(
	subscriptionRepository repository.SubscriptionRepository,
	userRepository repository.UserRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.SubscriptionService {
	return &serv{
		subscriptionRepository: subscriptionRepository,
		userRepository:         userRepository,
		logger:                 logger,
		cache:                  cache,
	}
}
