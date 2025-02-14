package payment

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
	"github.com/gorilla/websocket"
)

// @Summary Check payment status via WebSocket
// @Description Establishes a WebSocket connection to check the payment status for a given Telegram ID.
// @Tags Payment
// @Accept json
// @Produce json
// @Param telegramID path string true "Telegram ID of the user"
// @Success 200 {object} object '{"result": true}'
// @Failure 400 {object} payment.ErrorResponse "Invalid request parameters"
// @Failure 500 {object} payment.ErrorResponse "Internal server error"
// @Router /ws/v1/payment/check/{telegramID} [get]
func (wsh *WebSocketHandler) check(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		wsh.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, payment.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	fullURL, err := url.JoinPath(wsh.websocket.Payment.PaymentWsURL, telegramID)
	if err != nil {
		wsh.logger.Error("failed to build url", "error", err)
		c.JSON(http.StatusBadRequest, payment.ErrorResponse{
			Error:  "failed to join path ws payment connect and telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	u := websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}

	conn, err := u.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		wsh.logger.Error("failed to upgrade websocket server connection", "error", err)
		return
	}
	defer wsh.closeConnection(conn)

	if err := wsh.connectToPaymentService(c.Request.Context(), fullURL, conn); err != nil {
		wsh.logger.Error("failed to connect to payment service", "error", err)
		return
	}
}

// connectToPaymentService connect to payment service.
func (wsh *WebSocketHandler) connectToPaymentService(ctx context.Context, url string, clientConn *websocket.Conn) error {
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	defer wsh.closeConnection(conn)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, p, err := conn.ReadMessage()
			if err != nil {
				return err
			}

			if string(p) == `"{\"result\": true}"` {
				if err := clientConn.WriteMessage(websocket.TextMessage, []byte(`{"result": true}`)); err != nil {
					wsh.logger.Error("failed to send receive response to client", "error", err)
					return err
				}
				return nil
			} else {
				if err := clientConn.WriteMessage(websocket.TextMessage, []byte(`{"result": false}`)); err != nil {
					wsh.logger.Error("failed to send receive response to client", "error", err)
					return err
				}
			}
		}
	}
}

// closeConnection close websocket connection.
func (wsh *WebSocketHandler) closeConnection(conn *websocket.Conn) {
	if err := conn.Close(); err != nil {
		wsh.logger.Error("failed to close websocket connection", "error", err)
	}
}
