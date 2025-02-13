package recipeofdays

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
)

type Cron struct {
	recipeOfDaysService service.RecipeOfDaysService
	logger              *logger.Logger
	sleepDuration       int
	timeout             int
	maxErrorCount       int
	errorCount          int
}

func NewCron(
	ctx context.Context,
	recipeOfDaysService service.RecipeOfDaysService,
	worker config.WorkerConfig,
	logger *logger.Logger,
) *Cron {
	c := &Cron{
		recipeOfDaysService: recipeOfDaysService,
		logger:              logger,
		sleepDuration:       worker.LifeHackOfTheDay.SleepDuration,
		timeout:             worker.LifeHackOfTheDay.Timeout,
		maxErrorCount:       worker.LifeHackOfTheDay.MaxErrorCount,
		errorCount:          0,
	}

	go c.Start(ctx)

	return c
}

func (c *Cron) Start(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			c.logger.Info("cron recipe of days job stopped", slog.String("reason", ctx.Err().Error()))
			return
		}

		ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(c.timeout)*time.Second)
		err := c.recipeOfDaysService.Create(ctxTimeout)
		cancel()

		if err != nil {
			c.logger.Warn("error executing recipe of days service", slog.String("err", err.Error()))

			c.errorCount++

			if c.errorCount >= c.maxErrorCount {
				c.logger.Error("maximum error count reached in executing recipe of days service", slog.Int("error_count", c.errorCount))
				c.errorCount = 0
			}
		}

		time.Sleep(time.Duration(c.sleepDuration) * time.Minute)
	}
}
