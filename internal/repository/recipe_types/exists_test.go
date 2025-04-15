package recipetypes

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExists(t *testing.T) {
	type in struct {
		ctx   context.Context
		title string
	}

	type want struct {
		exists bool
		err    error
	}

	var (
		ctx          = context.TODO()
		title        = gofakeit.BookTitle()
		exists       = gofakeit.Bool()
		queryTimeout = int64(2)
	)

	tests := []struct {
		name               string
		mockPoolBehavior   func(m *poolsmocks.IPool, row *poolsmocks.RowMock)
		mockLoggerBehavior func(m *loggermocks.ILogger)
		in                 in
		want               want
	}{
		{
			name: "ok",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					title,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*bool"),
				).Run(func(args mock.Arguments) {
					ie := args.Get(0).(*bool)
					*ie = exists
				}).Return(nil)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[check a recipe type exists] execute repository")
			},
			in: in{
				ctx:   ctx,
				title: title,
			},
			want: want{
				exists: exists,
				err:    nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					title,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*bool"),
				).Run(func(args mock.Arguments) {
					ie := args.Get(0).(*bool)
					*ie = exists
				}).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[check a recipe type exists] execute repository")
				m.On("Error", "request timed out while check exists recipe type", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx:   ctx,
				title: title,
			},
			want: want{
				exists: false,
				err:    errors.New("the request timed out"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.RowMock) {
				m.On("QueryRow",
					mock.Anything,
					mock.Anything,
					title,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*bool"),
				).Run(func(args mock.Arguments) {
					ie := args.Get(0).(*bool)
					*ie = exists
				}).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[check a recipe type exists] execute repository")
				m.On("Error", "failed to check exists recipe type", "err", errors.New("database error"))
			},
			in: in{
				ctx:   ctx,
				title: title,
			},
			want: want{
				exists: false,
				err:    errors.New("could not check exists recipe type: database error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := poolsmocks.NewIPool(t)
			mockRow := poolsmocks.NewMockRow(t)
			mockLogger := loggermocks.NewILogger(t)

			if test.mockPoolBehavior != nil {
				test.mockPoolBehavior(mockPool, mockRow)
			}
			if test.mockLoggerBehavior != nil {
				test.mockLoggerBehavior(mockLogger)
			}

			pg := &postgres.Postgres{
				Pool:         mockPool,
				QueryTimeout: queryTimeout,
			}

			repository := NewRepository(mockLogger, pg)

			result, err := repository.Exists(test.in.ctx, test.in.title)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.exists, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
