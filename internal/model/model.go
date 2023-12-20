package model

import (
	"time"
)

type Event struct {
	EventType  string    `json:"event_type"`
	ScreenName string    `json:"screen_name"`
	Action     string    `json:"action" `
	EventTime  time.Time `json:"event_time"`
}
