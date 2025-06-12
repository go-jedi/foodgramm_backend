package clientassets

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	clientassets "github.com/go-jedi/foodgramm_backend/internal/domain/client_assets"
	fileserver "github.com/go-jedi/foodgramm_backend/internal/domain/file_server"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Upload client asset
// @Description Upload a new client asset file (image) to the server
// @Tags Client Assets
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param file formData file true "Image file to upload (supported types: PNG, JPEG)"
// @Success 201 {object} clientassets.ClientAssets "Successfully uploaded file, returns file information"
// @Failure 400 {object} clientassets.ErrorResponse "Bad request if no file provided or unsupported file type"
// @Failure 500 {object} clientassets.ErrorResponse "Internal server error if file processing fails"
// @Router /v1/client_assets [post]
func (h *Handler) create(c *gin.Context) {
	h.logger.Debug("[create a client assets] execute handler")

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("failed to get the first file for the provided form key", "error", err)
		c.JSON(http.StatusBadRequest, clientassets.ErrorResponse{
			Error:  "failed to get the first file for the provided form key",
			Detail: err.Error(),
		})
		return
	}

	contentType := file.Header.Get("Content-Type")
	if _, ok := fileserver.SupportedImageTypes[contentType]; !ok {
		h.logger.Error("unsupported file type: %s", contentType, "error")
		c.JSON(http.StatusBadRequest, clientassets.ErrorResponse{
			Error:  "unsupported file type",
			Detail: fmt.Errorf("%w: %s", apperrors.ErrUnsupportedFormat, contentType).Error(),
		})
		return
	}

	result, err := h.clientAssetsService.Create(c.Request.Context(), file)
	if err != nil {
		h.logger.Error("failed to create a client assets", "error", err)
		c.JSON(http.StatusInternalServerError, clientassets.ErrorResponse{
			Error:  "failed to create a client assets",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
