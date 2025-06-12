package dependencies

import (
	"github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/subscription"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	subscriptionRepository "github.com/go-jedi/foodgramm_backend/internal/repository/subscription"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	subscriptionService "github.com/go-jedi/foodgramm_backend/internal/service/subscription"
)

func (d *Dependencies) SubscriptionRepository() repository.SubscriptionRepository {
	if d.subscriptionRepository == nil {
		d.subscriptionRepository = subscriptionRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.subscriptionRepository
}

func (d *Dependencies) SubscriptionService() service.SubscriptionService {
	if d.subscriptionService == nil {
		d.subscriptionService = subscriptionService.NewService(
			d.SubscriptionRepository(),
			d.UserRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.subscriptionService
}

func (d *Dependencies) SubscriptionHandler() *subscription.Handler {
	if d.subscriptionHandler == nil {
		d.subscriptionHandler = subscription.NewHandler(
			d.SubscriptionService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.subscriptionHandler
}
