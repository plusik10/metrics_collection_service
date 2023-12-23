package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Event struct {
	EventType  string    `json:"event_type,omitempty"`
	ScreenName string    `json:"screen_name,omitempty"`
	Action     string    `json:"action,omitempty" `
	EventTime  time.Time `json:"event_time,omitempty"`
}

func (e Event) Validate() error {
	fmt.Println(e)

	return validation.ValidateStruct(&e,
		validation.Field(&e.EventType, validation.Required),
		validation.Field(&e.ScreenName, validation.Required),
		validation.Field(&e.Action, validation.Required),
		validation.Field(&e.EventTime, validation.Required),
	)
}
