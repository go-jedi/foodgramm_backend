package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	loggermocks "github.com/go-jedi/foodgrammm-backend/pkg/logger/mocks"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	poolsmocks "github.com/go-jedi/foodgrammm-backend/pkg/postgres/mocks"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	type in struct {
		ctx context.Context
		dto user.CreateDTO
	}

	type want struct {
		user user.User
		err  error
	}

	var (
		ctx        = context.TODO()
		telegramID = gofakeit.UUID()
		username   = gofakeit.Username()
		firstname  = gofakeit.FirstName()
		lastname   = gofakeit.LastName()
		createdAt  = time.Now()
		updatedAt  = time.Now()
		dto        = user.CreateDTO{
			TelegramID: telegramID,
			Username:   username,
			FirstName:  firstname,
			LastName:   lastname,
		}
		testUser = user.User{
			ID:         gofakeit.Int64(),
			TelegramID: telegramID,
			Username:   username,
			FirstName:  firstname,
			LastName:   lastname,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		}
		queryTimeout = int64(2)
	)

	tests := []struct {
		name               string
		mockPoolBehavior   func(m *poolsmocks.IPool, row *poolsmocks.MockRow)
		mockLoggerBehavior func(m *loggermocks.ILogger)
		in                 in
		want               want
	}{
		{
			name: "ok",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.MockRow) {
				rawData, _ := jsoniter.Marshal(dto)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Run(func(args mock.Arguments) {
					id := args.Get(0).(*int64)
					*id = testUser.ID

					tgID := args.Get(1).(*string)
					*tgID = testUser.TelegramID

					un := args.Get(2).(*string)
					*un = testUser.Username

					fn := args.Get(3).(*string)
					*fn = testUser.FirstName

					ln := args.Get(4).(*string)
					*ln = testUser.LastName

					ca := args.Get(5).(*time.Time)
					*ca = testUser.CreatedAt

					ua := args.Get(6).(*time.Time)
					*ua = testUser.UpdatedAt
				}).Return(nil)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.user_create($1);`,
					rawData,
				).Return(row)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create a new user] execute repository")
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				user: testUser,
				err:  nil,
			},
		},
		{
			name: "timeout error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.MockRow) {
				rawData, _ := jsoniter.Marshal(dto)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.user_create($1);`,
					rawData,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Return(context.DeadlineExceeded)
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create a new user] execute repository")
				m.On("Error", "request timed out while creating the user", "err", context.DeadlineExceeded)
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				user: user.User{},
				err:  errors.New("the request timed out"),
			},
		},
		{
			name: "database error",
			mockPoolBehavior: func(m *poolsmocks.IPool, row *poolsmocks.MockRow) {
				rawData, _ := jsoniter.Marshal(dto)

				m.On("QueryRow",
					mock.Anything,
					`SELECT * FROM public.user_create($1);`,
					rawData,
				).Return(row)

				row.On("Scan",
					mock.AnythingOfType("*int64"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*string"),
					mock.AnythingOfType("*time.Time"),
					mock.AnythingOfType("*time.Time"),
				).Return(errors.New("database error"))
			},
			mockLoggerBehavior: func(m *loggermocks.ILogger) {
				m.On("Debug", "[create a new user] execute repository")
				m.On("Error", "failed to create user", "err", errors.New("database error"))
			},
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				user: user.User{},
				err:  errors.New("could not create user"),
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

			result, err := repository.Create(test.in.ctx, test.in.dto)

			if test.want.err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.user, result)

			mockPool.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
