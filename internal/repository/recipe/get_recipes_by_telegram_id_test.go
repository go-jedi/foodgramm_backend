package recipe

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRecipesByTelegramID(t *testing.T) {
	type in struct {
		ctx        context.Context
		telegramID string
	}

	type want struct {
		recipes []recipe.Recipes
		err     error
	}

	var (
		ctx         = context.TODO()
		telegramID  = gofakeit.UUID()
		testRecipes = []recipe.Recipes{
			{
				ID:         gofakeit.Int64(),
				TelegramID: telegramID,
				Title:      gofakeit.BookTitle(),
				Content: [][]recipe.Content{
					{
						{
							ID:                gofakeit.Int64(),
							Type:              gofakeit.BookTitle(),
							Title:             gofakeit.BookTitle(),
							RecipePreparation: gofakeit.BookTitle(),
							Calories:          gofakeit.BookTitle(),
							Bzhu:              gofakeit.BookTitle(),
							Ingredients:       gofakeit.NiceColors(),
							MethodPreparation: gofakeit.NiceColors(),
						},
					},
				},
				CreatedAt: gofakeit.Date(),
				UpdatedAt: gofakeit.Date(),
			},
		}
		queryTimeout = int64(2)
	)

	tests := []struct {
		name               string
		mockPoolBehavior   func(m *poolsmocks.IPool)
		mockLoggerBehavior func(m *loggermocks.ILogger)
		in                 in
		want               want
	}{
		{
			name: "ok",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return len(testRecipes) > 0
					},
					ScanFunc: func(dest ...any) error {
						if len(testRecipes) == 0 {
							return pgx.ErrNoRows
						}

						tr := testRecipes[0]
						testRecipes = testRecipes[1:]

						*dest[0].(*int64) = tr.ID
						*dest[1].(*string) = tr.TelegramID
						*dest[2].(*string) = tr.Title
						*dest[3].(*[][]recipe.Content) = tr.Content
						*dest[4].(*time.Time) = tr.CreatedAt
						*dest[5].(*time.Time) = tr.UpdatedAt

						return nil
					},
					ErrFunc: func() error {
						return nil
					},
					CloseFunc: func() {},
					CommandTagFunc: func() pgconn.CommandTag {
						return pgconn.NewCommandTag("SELECT")
					},
					FieldDescriptionsFunc: func() []pgconn.FieldDescription {
						return []pgconn.FieldDescription{
							{Name: "id", DataTypeOID: 20},
							{Name: "telegram_id", DataTypeOID: 20},
							{Name: "title", DataTypeOID: 25},
							{Name: "content", DataTypeOID: 25},
							{Name: "created_at", DataTypeOID: 1114},
							{Name: "updated_at", DataTypeOID: 1114},
						}
					},
					ValuesFunc: func() ([]any, error) {
						if len(testRecipes) == 0 {
							return nil, pgx.ErrNoRows
						}
						tr := testRecipes[0]
						return []any{
							tr.ID,
							tr.TelegramID,
							tr.Title,
							tr.Content,
							tr.CreatedAt,
							tr.UpdatedAt,
						}, nil
					},
					RawValuesFunc: func() [][]byte {
						return nil
					},
				}
				m.On("Query", mock.Anything, mock.Anything, telegramID).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipes by telegram id] execute repository", mock.Anything)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				recipes: testRecipes,
				err:     nil,
			},
		},
		{
			name: "query error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything, telegramID).Return(nil, pgx.ErrNoRows)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipes by telegram id] execute repository", mock.Anything)
				m.On("Error", "failed to get recipes by telegram id", "err", pgx.ErrNoRows)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				recipes: nil,
				err:     fmt.Errorf("could not get recipes by telegram id: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "context timeout",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything, telegramID).Return(nil, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipes by telegram id] execute repository", mock.Anything)
				m.On("Error", "request timed out while get recipes by telegram id", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				recipes: nil,
				err:     fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
			},
		},
		{
			name: "scan error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return true
					},
					ScanFunc: func(dest ...any) error {
						return pgx.ErrNoRows
					},
					ErrFunc: func() error {
						return nil
					},
					CloseFunc: func() {},
				}
				m.On("Query", mock.Anything, mock.Anything, telegramID).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipes by telegram id] execute repository", mock.Anything)
				m.On("Error", "failed to scan row to get recipes by telegram id", "err", pgx.ErrNoRows)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				recipes: nil,
				err:     fmt.Errorf("failed to scan row to get recipes by telegram id: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "rows error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				rows := &poolsmocks.RowsMock{
					NextFunc: func() bool {
						return false
					},
					ScanFunc: func(dest ...any) error {
						return nil
					},
					ErrFunc: func() error {
						return pgx.ErrNoRows
					},
					CloseFunc: func() {},
				}
				m.On("Query", mock.Anything, mock.Anything, telegramID).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get recipes by telegram id] execute repository", mock.Anything)
				m.On("Error", "failed to get recipes by telegram id", "err", pgx.ErrNoRows)
			},
			in: in{
				ctx:        ctx,
				telegramID: telegramID,
			},
			want: want{
				recipes: nil,
				err:     fmt.Errorf("failed to get recipes by telegram id: %w", pgx.ErrNoRows),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := poolsmocks.NewIPool(t)
			mockLogger := loggermocks.NewILogger(t)

			if test.mockPoolBehavior != nil {
				test.mockPoolBehavior(mockPool)
			}
			if test.mockLoggerBehavior != nil {
				test.mockLoggerBehavior(mockLogger)
			}

			pg := &postgres.Postgres{
				Pool:         mockPool,
				QueryTimeout: queryTimeout,
			}

			repository := NewRepository(mockLogger, pg)

			result, err := repository.GetRecipesByTelegramID(test.in.ctx, test.in.telegramID)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipes, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
