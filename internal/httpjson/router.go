package httpjson

import (
	"net/http"

	"github.com/hferr/events-api/internal/app"
)

type Handler struct {
	eventSvc app.EventService
}

func NewHandler(eventSvs app.EventService) *Handler {
	return &Handler{
		eventSvc: eventSvs,
	}
}

func (h *Handler) NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("POST /events", h.CreateEvent)
	mux.HandleFunc("GET /events", h.ListEvents)
	mux.HandleFunc("GET /events/{id}", h.GetEventByID)

	return mux
}
