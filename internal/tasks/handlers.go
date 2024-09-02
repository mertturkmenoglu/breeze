package tasks

import (
	"breeze/internal/db"
	"context"
	"net/http"

	"github.com/hibiken/asynq"
)

func (ts *Tasks) HandleCheckStatusTask(_ context.Context, t *asynq.Task) error {
	p, err := parse[CheckStatusPayload](t.Payload())

	if err != nil {
		return err
	}

	ts.logger.Info("Handling task", ts.logger.Args("task", "check status", "id", p.ID, "url", p.URL))

	res, err := http.Get(p.URL)

	if err != nil {
		ts.logger.Error("error getting page", ts.logger.Args("err", err))
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		ts.logger.Info("Page is up", ts.logger.Args("id", p.ID, "url", p.URL))

		return ts.db.Queries.UpdatePageStatus(context.Background(), db.UpdatePageStatusParams{
			ID:     p.ID,
			Status: db.PagestatusONLINE,
		})
	}

	ts.logger.Info("Page is down", ts.logger.Args("id", p.ID, "url", p.URL))

	return ts.db.Queries.UpdatePageStatus(context.Background(), db.UpdatePageStatusParams{
		ID:     p.ID,
		Status: db.PagestatusOFFLINE,
	})
}
