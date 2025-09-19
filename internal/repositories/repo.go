package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/hferr/events-api/internal/app"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

const InsertEventQuery = `
	INSERT INTO events 
		(id, title, description, start_time, end_time, created_at)
	VALUES
		($1, $2, $3, $4, $5, NOW())
	RETURNING *
`

func (r Repo) CreateEvent(ctx context.Context, event app.Event) (app.Event, error) {
	var createdEvent app.Event
	err := r.db.QueryRowContext(
		ctx,
		InsertEventQuery,
		event.Id,
		event.Title,
		event.Description,
		event.StartTime,
		event.EndTime,
	).Scan(
		&createdEvent.Id,
		&createdEvent.Title,
		&createdEvent.Description,
		&createdEvent.StartTime,
		&createdEvent.EndTime,
		&createdEvent.CreatedAt,
	)
	if err != nil {
		return app.Event{}, err
	}

	return createdEvent, nil
}

const ListEventsQuery = `SELECT * FROM events ORDER BY start_time ASC`

func (r Repo) ListEvents(ctx context.Context) ([]app.Event, error) {
	var events []app.Event
	rows, err := r.db.QueryContext(ctx, ListEventsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var eventRow app.Event
		err := rows.Scan(
			&eventRow.Id,
			&eventRow.Title,
			&eventRow.Description,
			&eventRow.StartTime,
			&eventRow.EndTime,
			&eventRow.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, eventRow)
	}

	return events, nil
}

const GetEventByIdQuery = `SELECT * FROM events WHERE id=$1`

func (r Repo) GetEventByID(ctx context.Context, id uuid.UUID) (*app.Event, error) {
	var event app.Event
	err := r.db.QueryRowContext(
		ctx,
		GetEventByIdQuery,
		id,
	).Scan(
		&event.Id,
		&event.Title,
		&event.Description,
		&event.StartTime,
		&event.EndTime,
		&event.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}
