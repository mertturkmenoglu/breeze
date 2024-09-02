package cron

import (
	"breeze/internal/db"
	"breeze/internal/tasks"
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/pterm/pterm"
)

func New() gocron.Scheduler {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic(err.Error())
	}

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	db := db.NewDb()
	ts := tasks.New(db)

	cronstr := "*/1 * * * *"
	_, err = scheduler.NewJob(gocron.CronJob(cronstr, false), gocron.NewTask(func() {
		logger.Info("Running cron job", logger.Args("job", "check status"))
		res, err := db.Queries.GetPagesThatNeedChecking(context.Background())

		if err != nil {
			logger.Error("error getting pages", logger.Args("err", err))
			return
		}

		logger.Info("Running cron job, got pages need checking", logger.Args("# of pages", len(res)))

		for _, page := range res {
			_, err := ts.CreateAndEnqueue(tasks.TypeCheckStatus, tasks.CheckStatusPayload{
				ID:  page.ID,
				URL: page.Url,
			})

			if err != nil {
				logger.Error("error creating and enqueuing task", logger.Args("err", err))
			}
		}
	}))

	if err != nil {
		panic(err.Error())
	}

	return scheduler
}
