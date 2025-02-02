package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Get Exclude Products by User ID
// @Description Get excluded products for a user by their unique User ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param userID path int64 true "User ID"
// @Success 200 {object} product.UserExcludedProducts "Excluded products for the user"
// @Failure 400 {object} product.ErrorResponse
// @Failure 500 {object} product.ErrorResponse
// @Router /v1/product/exclude/user/{userID} [get]
func (h *Handler) getExcludeProductsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		h.logger.Error("failed to get param userID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to get param userID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	userIDInt, err := strconv.ParseInt(userID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed parse string to int64",
			Detail: err.Error(),
		})
		return
	}

	if userIDInt <= 0 {
		h.logger.Error("userID is less than or equal to zero", "error", err)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "userID must be greater than zero",
			Detail: "the provided userID is less than or equal to zero, which is not allowed",
		})
		return
	}

	result, err := h.productService.GetExcludeProductsByUserID(c.Request.Context(), userIDInt)
	if err != nil {
		h.logger.Error("failed to get exclude products by user id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to get exclude products by user id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
