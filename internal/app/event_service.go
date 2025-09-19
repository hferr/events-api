package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrEventValidation = errors.New("event validation error")

type EventService interface {
	CreateEvent(ctx context.Context, event Event) (Event, error)
	ListEvents(ctx context.Context) ([]Event, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (*Event, error)
}

type eventService struct {
	repo Repo
}

func NewEventService(repo Repo) EventService {
	return eventService{
		repo: repo,
	}
}

func (s eventService) CreateEvent(ctx context.Context, event Event) (Event, error) {
	if err := s.validateEvent(event); err != nil {
		return Event{}, err
	}

	event.Id = uuid.New()

	return s.repo.CreateEvent(ctx, event)
}

func (s eventService) ListEvents(ctx context.Context) ([]Event, error) {
	return s.repo.ListEvents(ctx)
}

func (s eventService) GetEventByID(ctx context.Context, id uuid.UUID) (*Event, error) {
	return s.repo.GetEventByID(ctx, id)
}

func (s eventService) validateEvent(event Event) error {
	var errs []error

	if event.Title == "" {
		errs = append(errs, fmt.Errorf("title cannot be empty"))
	}

	if len(event.Title) > 100 {
		errs = append(errs, fmt.Errorf("title must have less than 100 characters"))
	}

	if event.StartTime.After(event.EndTime) {
		errs = append(errs, fmt.Errorf("start time must be before end time"))
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w: %w", ErrEventValidation, errors.Join(errs...))
	}

	return nil
}
