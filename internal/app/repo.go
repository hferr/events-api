package app

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	CreateEvent(ctx context.Context, event Event) (Event, error)
	ListEvents(ctx context.Context) ([]Event, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (*Event, error)
}
