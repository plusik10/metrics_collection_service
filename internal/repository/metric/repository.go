package metric

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/plusik10/metrics_collection_service/internal/model"
	"github.com/plusik10/metrics_collection_service/internal/repository"
	"github.com/plusik10/metrics_collection_service/pkg/db"
)

const (
	TableName      = "event"
	EventType      = "event_type"
	EventTimestamp = "event_timestamp"
	ScreenName     = "screen_name"
	Action         = "action"
)

type metricRepository struct {
	client db.Client
}

func NewMetricRepository(db db.Client) repository.Repository {
	return &metricRepository{
		client: db,
	}
}

func (m *metricRepository) Create(ctx context.Context, event *model.Event) error {
	query, args, err := squirrel.Insert(TableName).
		Columns(EventType, EventTimestamp, ScreenName, Action).PlaceholderFormat(squirrel.Dollar).
		Values(event.EventType, event.EventTime, event.ScreenName, event.Action).Suffix("returning id").ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Create",
		QueryRow: query,
	}

	_, err = m.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil

}
