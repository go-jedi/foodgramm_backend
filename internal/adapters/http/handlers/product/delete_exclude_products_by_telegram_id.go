package product

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary      Delete excluded products by Telegram ID
// @Description  This endpoint allows you to delete excluded products for a user identified by their Telegram ID.
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Authorization token" default(Bearer <token>)
// @Param        telegramID path string true "Telegram ID of the user"
// @Param        product query string true "Product name to be deleted"
// @Success      200 {object} UserExcludedProducts "Successfully deleted excluded product"
// @Failure      400 {object} product.ErrorResponse "Bad Request"
// @Failure      500 {object} product.ErrorResponse "Internal Server Error"
// @Router       /v1/product/exclude/telegram/{telegramID} [delete]
func (h *Handler) deleteExcludeProductsByTelegramID(c *gin.Context) {
	h.logger.Debug("[deleteExcludeProductsByTelegramID] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	prod := c.Query("product")
	if prod == "" {
		h.logger.Error("failed to get query product", "error", apperrors.ErrQueryIsRequired)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to get query product",
			Detail: apperrors.ErrQueryIsRequired.Error(),
		})
		return
	}

	result, err := h.productService.DeleteExcludeProductsByTelegramID(c.Request.Context(), telegramID, prod)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, product.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to delete exclude product by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to delete exclude product by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
