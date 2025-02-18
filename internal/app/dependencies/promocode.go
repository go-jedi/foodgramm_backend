package dependencies

import (
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/promocode"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	promoCodeRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/promocode"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	promoCodeService "github.com/go-jedi/foodgrammm-backend/internal/service/promocode"
)

func (d *Dependencies) PromoCodeRepository() repository.PromoCodeRepository {
	if d.promoCodeRepository == nil {
		d.promoCodeRepository = promoCodeRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.promoCodeRepository
}

func (d *Dependencies) PromoCodeService() service.PromoCodeService {
	if d.promoCodeService == nil {
		d.promoCodeService = promoCodeService.NewService(
			d.PromoCodeRepository(),
			d.UserRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.promoCodeService
}

func (d *Dependencies) PromoCodeHandler() *promocode.Handler {
	if d.promoCodeHandler == nil {
		d.promoCodeHandler = promocode.NewHandler(
			d.PromoCodeService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.promoCodeHandler
}
