package recipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary      Get free recipes by Telegram ID
// @Description  This endpoint retrieves the free recipes information for a user identified by their Telegram ID.
// @Tags         Recipe
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token" default(Bearer <token>)
// @Param        telegramID path string true "Telegram ID of the user"
// @Success      200 {object} UserFreeRecipes "Successfully retrieved free recipes information"
// @Failure      400 {object} recipe.ErrorResponse "Bad Request"
// @Failure      500 {object} recipe.ErrorResponse "Internal Server Error"
// @Router       /v1/recipe/free/telegram/{telegramID} [get]
func (h *Handler) getFreeRecipesByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.recipeService.GetFreeRecipesByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, recipe.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to get recipes by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to get recipes by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
