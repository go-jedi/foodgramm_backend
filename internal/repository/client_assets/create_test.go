package clientassets

/*
import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	clientassets "github.com/go-jedi/foodgrammm-backend/internal/domain/client_assets"
	fileserver "github.com/go-jedi/foodgrammm-backend/internal/domain/file_server"
	fileservermocks "github.com/go-jedi/foodgrammm-backend/pkg/file_server/mocks"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
)

func TestCreate(t *testing.T) {
	type in struct {
		ctx  context.Context
		data fileserver.UploadAndConvertToWebpResponse
	}

	type want struct {
		clientAssets clientassets.ClientAssets
		err          error
	}

	var (
		ctx            = context.TODO()
		nameFile       = gofakeit.Name()
		serverPathFile = gofakeit.URL()
		clientPathFile = gofakeit.URL()
		extension      = gofakeit.FileExtension()
		quality        = gofakeit.IntRange(1, 100)
		oldNameFile    = gofakeit.Name()
		oldExtension   = gofakeit.FileExtension()
		data           = fileserver.UploadAndConvertToWebpResponse{
			NameFile:       nameFile,
			ServerPathFile: serverPathFile,
			ClientPathFile: clientPathFile,
			Extension:      extension,
			Quality:        quality,
			OldNameFile:    oldNameFile,
			OldExtension:   oldExtension,
		}
		testClientAssets = clientassets.ClientAssets{
			ID:             gofakeit.Int64(),
			NameFile:       nameFile,
			ServerPathFile: serverPathFile,
			ClientPathFile: clientPathFile,
			Extension:      extension,
			Quality:        quality,
			OldNameFile:    oldNameFile,
			OldExtension:   oldExtension,
			CreatedAt:      gofakeit.Date(),
			UpdatedAt:      gofakeit.Date(),
		}
		queryTimeout = int64(2)
	)

	tests := []struct {
		name                   string
		mockPoolBehavior       func(m *poolsmocks.IPool, row *poolsmocks.RowMock)
		mockLoggerBehavior     func(m *loggermocks.ILogger)
		mockFileServerBehavior func(m *fileservermocks.IFileServer)
		in                     in
		want                   want
	}{
		{},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := poolsmocks.NewIPool(t)
			mockRow := poolsmocks.NewMockRow(t)
			mockLogger := loggermocks.NewILogger(t)
			mockFileServer := fileservermocks.NewIFileServer(t)

			if test.mockPoolBehavior != nil {
				test.mockPoolBehavior(mockPool, mockRow)
			}
			if test.mockLoggerBehavior != nil {
				test.mockLoggerBehavior(mockLogger)
			}
			if test.mockFileServerBehavior != nil {
				//
			}

			pg := &postgres.Postgres{
				Pool:         mockPool,
				QueryTimeout: queryTimeout,
			}
		})
	}
}
*/
