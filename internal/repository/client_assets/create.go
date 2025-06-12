package clientassets

import (
	"context"
	"errors"
	"fmt"
	"time"

	clientassets "github.com/go-jedi/foodgramm_backend/internal/domain/client_assets"
	fileserver "github.com/go-jedi/foodgramm_backend/internal/domain/file_server"
)

func (r *repo) Create(ctx context.Context, data fileserver.UploadAndConvertToWebpResponse) (clientassets.ClientAssets, error) {
	r.logger.Debug("[create a client assets] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO client_assets(
			name_file,
		    server_path_file,
		    client_path_file,
		    extension,
		    quality,
		    old_name_file,
		    old_extension
		) VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`

	var ca clientassets.ClientAssets

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		data.NameFile, data.ServerPathFile, data.ClientPathFile,
		data.Extension, data.Quality, data.OldNameFile, data.OldExtension,
	).Scan(
		&ca.ID, &ca.NameFile, &ca.ServerPathFile,
		&ca.ClientPathFile, &ca.Extension, &ca.Quality,
		&ca.OldNameFile, &ca.OldExtension, &ca.CreatedAt, &ca.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while create a client assets", "err", err)
			return clientassets.ClientAssets{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create a client assets", "err", err)
		return clientassets.ClientAssets{}, fmt.Errorf("could not create a client assets: %w", err)
	}

	return ca, nil
}
