package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (ws *WebSocket) check(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		ws.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, payment.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}
}
