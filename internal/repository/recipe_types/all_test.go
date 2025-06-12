package recipetypes

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
	loggermocks "github.com/go-jedi/foodgramm_backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgramm_backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAll(t *testing.T) {
	type in struct {
		ctx context.Context
	}

	type want struct {
		recipeTypes []recipetypes.RecipeTypes
		err         error
	}

	var (
		ctx             = context.TODO()
		testRecipeTypes = []recipetypes.RecipeTypes{
			{
				ID:        gofakeit.Int64(),
				Title:     gofakeit.BookTitle(),
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
						return len(testRecipeTypes) > 0
					},
					ScanFunc: func(dest ...any) error {
						if len(testRecipeTypes) == 0 {
							return pgx.ErrNoRows
						}

						rt := testRecipeTypes[0]
						testRecipeTypes = testRecipeTypes[1:]

						*dest[0].(*int64) = rt.ID
						*dest[1].(*string) = rt.Title
						*dest[2].(*time.Time) = rt.CreatedAt
						*dest[3].(*time.Time) = rt.UpdatedAt

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
							{Name: "title", DataTypeOID: 20},
							{Name: "created_at", DataTypeOID: 1114},
							{Name: "updated_at", DataTypeOID: 1114},
						}
					},
					ValuesFunc: func() ([]any, error) {
						if len(testRecipeTypes) == 0 {
							return nil, pgx.ErrNoRows
						}
						rt := testRecipeTypes[0]
						return []any{
							rt.ID,
							rt.Title,
							rt.CreatedAt,
							rt.UpdatedAt,
						}, nil
					},
					RawValuesFunc: func() [][]byte {
						return nil
					},
				}
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipe types] execute repository", mock.Anything)
			},
			in: in{ctx: ctx},
			want: want{
				recipeTypes: testRecipeTypes,
				err:         nil,
			},
		},
		{
			name: "query error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, pgx.ErrNoRows)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipe types] execute repository", mock.Anything)
				m.On("Error", "failed to get recipe types", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipeTypes: nil,
				err:         fmt.Errorf("could not get recipe types: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "context timeout",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipe types] execute repository", mock.Anything)
				m.On("Error", "request timed out while get recipe types", "err", context.DeadlineExceeded)
			},
			in: in{ctx: ctx},
			want: want{
				recipeTypes: nil,
				err:         fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
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
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipe types] execute repository", mock.Anything)
				m.On("Error", "failed to scan row to get all recipe types", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipeTypes: nil,
				err:         fmt.Errorf("failed to scan row to get all recipe types: %w", pgx.ErrNoRows),
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
				m.On("Query", mock.Anything, mock.Anything).Return(rows, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipe types] execute repository", mock.Anything)
				m.On("Error", "failed to get all recipe types", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipeTypes: nil,
				err:         fmt.Errorf("failed to get all recipe types: %w", pgx.ErrNoRows),
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

			result, err := repository.All(test.in.ctx)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.recipeTypes, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
