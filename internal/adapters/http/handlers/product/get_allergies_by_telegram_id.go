package product

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Get Allergies by Telegram ID
// @Description Retrieve user allergies based on Telegram ID
// @Tags Product
// @Accept json
// @Produce json
// @Param        Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user"
// @Success 200 {object} product.UserExcludedProducts "User Excluded Products Response"
// @Failure 400 {object} product.ErrorResponse "Bad Request"
// @Failure 404 {object} product.ErrorResponse "User Not Found"
// @Failure 500 {object} product.ErrorResponse "Internal Server Error"
// @Router /v1/product/allergy/telegram/{telegramID} [get]
func (h *Handler) getAllergiesByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.productService.GetAllergiesByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, product.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to get allergies by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to get allergies by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
