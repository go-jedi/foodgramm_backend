package clientassets

import (
	"context"
	"mime/multipart"
	"os"

	clientassets "github.com/go-jedi/foodgramm_backend/internal/domain/client_assets"
)

func (s *serv) Create(ctx context.Context, file *multipart.FileHeader) (clientassets.ClientAssets, error) {
	s.logger.Debug("[create a client assets] execute service")

	// convert png or jpg image to webp and upload.
	imageData, err := s.fileServer.UploadAndConvertToWebP(ctx, file)
	if err != nil {
		return clientassets.ClientAssets{}, err
	}

	result, err := s.clientAssetsRepository.Create(ctx, imageData)
	if err != nil {
		// compensating action - delete the saved image.
		if err := os.Remove(imageData.ServerPathFile); err != nil {
			s.logger.Warn("failed to remove image after db error", "warn", err)
		}
		return clientassets.ClientAssets{}, err
	}

	return result, nil
}
