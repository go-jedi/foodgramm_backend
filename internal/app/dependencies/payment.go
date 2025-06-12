package dependencies

import (
	"github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/payment"
	paymentwebsocket "github.com/go-jedi/foodgramm_backend/internal/adapters/websocket/payment"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	paymentService "github.com/go-jedi/foodgramm_backend/internal/service/payment"
)

func (d *Dependencies) PaymentService() service.PaymentService {
	if d.paymentService == nil {
		d.paymentService = paymentService.NewService(
			d.UserRepository(),
			d.client,
			d.logger,
		)
	}

	return d.paymentService
}

func (d *Dependencies) PaymentHandler() *payment.Handler {
	if d.paymentHandler == nil {
		d.paymentHandler = payment.NewHandler(
			d.PaymentService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.paymentHandler
}

func (d *Dependencies) PaymentWebSocket() *paymentwebsocket.WebSocketHandler {
	if d.paymentWebSocketHandler == nil {
		d.paymentWebSocketHandler = paymentwebsocket.NewWebSocketHandler(
			d.SubscriptionService(),
			d.cookie,
			d.websocket,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.paymentWebSocketHandler
}
