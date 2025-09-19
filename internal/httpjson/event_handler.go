package httpjson

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hferr/events-api/internal/app"
)

type CreateEventRequest struct {
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "could not parse request body", http.StatusBadRequest)
		return
	}

	event := app.Event{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}

	createdEvent, err := h.eventSvc.CreateEvent(r.Context(), event)
	if err != nil {
		if errors.Is(err, app.ErrEventValidation) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(fmtValidationError(err))
			return
		}

		http.Error(w, "could not create event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEvent)
}

func (h *Handler) ListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.eventSvc.ListEvents(r.Context())
	if err != nil {
		http.Error(w, "could not list events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(events) > 0 {
		json.NewEncoder(w).Encode(events)
	} else {
		json.NewEncoder(w).Encode([]string{})
	}
}

func (h *Handler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "could not parse id from query", http.StatusBadRequest)
		return
	}

	event, err := h.eventSvc.GetEventByID(r.Context(), id)
	if err != nil {
		http.Error(w, "could not get event by id", http.StatusInternalServerError)
		return
	}

	if event == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func fmtValidationError(err error) string {
	if !errors.Is(err, app.ErrEventValidation) {
		return err.Error()
	}

	fmtError := err.Error()
	fmtError = strings.ReplaceAll(fmtError, "\n", "; ")

	return fmtError
}
