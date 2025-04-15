package recipeofdays

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	type in struct {
		ctx  context.Context
		data parser.ParsedRecipeOfDays
	}

	type want struct {
		err error
	}

	var (
		ctx  = context.TODO()
		data = parser.ParsedRecipeOfDays{
			Title: gofakeit.BookTitle(),
			Lifehack: parser.Lifehack{
				Name:        gofakeit.Name(),
				Description: gofakeit.ProductDescription(),
			},
			Content: [][]parser.Content{
				{
					{
						ID:                gofakeit.Int64(),
						Type:              gofakeit.AnimalType(),
						Title:             gofakeit.BookTitle(),
						RecipePreparation: gofakeit.ProductDescription(),
						Calories:          gofakeit.ProductDescription(),
						Bzhu:              gofakeit.ProductDescription(),
						Ingredients:       gofakeit.NiceColors(),
						MethodPreparation: gofakeit.NiceColors(),
					},
				},
			},
		}
		queryTimeout      = int64(2)
		successCommandTag = pgconn.NewCommandTag("INSERT 0 1")
		noRowsCommandTag  = pgconn.NewCommandTag("INSERT 0 0")
		emptyCommandTag   = pgconn.NewCommandTag("")
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
				m.On("Exec", mock.Anything, mock.Anything, data.Title, data.Lifehack, data.Content).
					Return(successCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe of the day] execute repository", mock.Anything)
			},
			in: in{
				ctx:  ctx,
				data: data,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, data.Title, data.Lifehack, data.Content).
					Return(emptyCommandTag, context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe of the day] execute repository", mock.Anything)
				m.On("Error", "request timed out while create recipe of days", "err", mock.Anything)
			},
			in: in{
				ctx:  ctx,
				data: data,
			},
			want: want{
				err: fmt.Errorf("the request timed out: %w", context.DeadlineExceeded),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, data.Title, data.Lifehack, data.Content).
					Return(emptyCommandTag, errors.New("some db error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe of the day] execute repository", mock.Anything)
				m.On("Error", "failed to create recipe of days", "err", mock.Anything)
			},
			in: in{
				ctx:  ctx,
				data: data,
			},
			want: want{
				err: errors.New("could not create recipe of days: some db error"),
			},
		},
		{
			name: "no rows affected",
			mockPoolBehavior: func(m *poolsmocks.IPool) {
				m.On("Exec", mock.Anything, mock.Anything, data.Title, data.Lifehack, data.Content).
					Return(noRowsCommandTag, nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create recipe of the day] execute repository", mock.Anything)
			},
			in: in{
				ctx:  ctx,
				data: data,
			},
			want: want{
				err: apperrors.ErrNoRowsWereAffected,
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

			err := repository.Create(test.in.ctx, test.in.data)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
