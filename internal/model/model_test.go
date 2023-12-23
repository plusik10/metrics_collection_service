package model

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestEvent_Validate_Positive(t *testing.T) {

	testCase := []struct {
		name  string
		event Event
	}{
		{
			name: "Success test",
			event: Event{
				EventType:  gofakeit.BeerName(),
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
			},
		},
		{
			name: "Empty first field",
			event: Event{
				EventType:  "12",
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
			},
		},
	}

	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.event.Validate()
			require.NoError(t, err)
		})
	}
}

func TestEvent_Validate_Nigative(t *testing.T) {

	testCase := []struct {
		name  string
		event Event
	}{
		{
			name: "Success test",
			event: Event{
				EventType:  gofakeit.BeerName(),
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
			},
		},
		{
			name: "Empty first field",
			event: Event{
				EventType:  "",
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
			},
		},
		{
			name: "Empty second field",
			event: Event{
				EventType: gofakeit.BeerName(),
				Action:    gofakeit.BeerName(),
				EventTime: gofakeit.Date(),
			},
		},
		{
			name: "Empty third field",
			event: Event{
				EventType:  gofakeit.BeerName(),
				ScreenName: gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
			},
		},
		{
			name: "Empty fourth field",
			event: Event{
				EventType:  gofakeit.BeerName(),
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
			},
		},
	}

	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.event.Validate()
			require.Error(t, err)
		})
	}

}
