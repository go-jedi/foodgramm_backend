package recipeevent

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
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
		recipesEvent []recipeevent.Recipe
		err          error
	}

	var (
		ctx              = context.TODO()
		testRecipesEvent = []recipeevent.Recipe{
			{
				ID:     gofakeit.Int64(),
				TypeID: gofakeit.Int64(),
				Title:  gofakeit.BookTitle(),
				Content: [][]recipeevent.Content{
					{
						{
							ID:                gofakeit.Int64(),
							Type:              gofakeit.UUID(),
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
						return len(testRecipesEvent) > 0
					},
					ScanFunc: func(dest ...any) error {
						if len(testRecipesEvent) == 0 {
							return pgx.ErrNoRows
						}

						rt := testRecipesEvent[0]
						testRecipesEvent = testRecipesEvent[1:]

						*dest[0].(*int64) = rt.ID
						*dest[1].(*int64) = rt.TypeID
						*dest[2].(*string) = rt.Title
						*dest[3].(*[][]recipeevent.Content) = rt.Content
						*dest[4].(*time.Time) = rt.CreatedAt
						*dest[5].(*time.Time) = rt.UpdatedAt

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
							{Name: "type_id", DataTypeOID: 20},
							{Name: "title", DataTypeOID: 20},
							{Name: "content", DataTypeOID: 20},
							{Name: "created_at", DataTypeOID: 1114},
							{Name: "updated_at", DataTypeOID: 1114},
						}
					},
					ValuesFunc: func() ([]any, error) {
						if len(testRecipesEvent) == 0 {
							return nil, pgx.ErrNoRows
						}
						rt := testRecipesEvent[0]
						return []any{
							rt.ID,
							rt.TypeID,
							rt.Title,
							rt.Content,
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
				m.On("Debug", "[get all recipes event] execute repository", mock.Anything)
			},
			in: in{ctx: ctx},
			want: want{
				recipesEvent: testRecipesEvent,
				err:          nil,
			},
		},
		{
			name: "query error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, pgx.ErrNoRows)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipes event] execute repository", mock.Anything)
				m.On("Error", "failed to get recipes event", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipesEvent: nil,
				err:          fmt.Errorf("could not get recipes event: %w", pgx.ErrNoRows),
			},
		},
		{
			name: "context timeout",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Query", mock.Anything, mock.Anything).Return(nil, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[get all recipes event] execute repository", mock.Anything)
				m.On("Error", "request timed out while get recipes event", "err", context.DeadlineExceeded)
			},
			in: in{ctx: ctx},
			want: want{
				recipesEvent: nil,
				err:          fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
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
				m.On("Debug", "[get all recipes event] execute repository", mock.Anything)
				m.On("Error", "failed to scan row to get all recipes event", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipesEvent: nil,
				err:          fmt.Errorf("failed to scan row to get all recipes event: %w", pgx.ErrNoRows),
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
				m.On("Debug", "[get all recipes event] execute repository", mock.Anything)
				m.On("Error", "failed to get all recipes event", "err", pgx.ErrNoRows)
			},
			in: in{ctx: ctx},
			want: want{
				recipesEvent: nil,
				err:          fmt.Errorf("failed to get all recipes event: %w", pgx.ErrNoRows),
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

			assert.Equal(t, test.want.recipesEvent, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
