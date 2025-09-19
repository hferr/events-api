package app

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}
