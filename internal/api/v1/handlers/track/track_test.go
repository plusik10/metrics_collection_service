package track

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/plusik10/metrics_collection_service/internal/model"
	"github.com/plusik10/metrics_collection_service/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewHandler_Nigative(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		event *model.Event
		OK    bool
		Err   string
	}{
		{
			event: &model.Event{
				EventType: "",
			},
			OK: false,
			Err: "action: cannot be blank; event_time: cannot be blank; event_type: " +
				"cannot be blank; screen_name: cannot be blank.",
		},
		{
			event: &model.Event{
				EventType:  "",
				EventTime:  gofakeit.Date(),
				ScreenName: "",
				Action:     "",
			},
			OK:  false,
			Err: "action: cannot be blank; event_type: cannot be blank; screen_name: cannot be blank.",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run("handler should", func(t *testing.T) {
			t.Parallel()
			metricMocks := mocks.NewMetricService(t)
			handler := New(ctx, metricMocks)
			b, err := json.Marshal(tc.event)
			assert.NoError(t, err)
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/track", bytes.NewReader(b))
			if err != nil {
				assert.Error(t, err)
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			var res ResponseTrack
			json.NewDecoder(rr.Body).Decode(&res)
			require.Equal(t, tc.OK, res.OK)
			require.Equal(t, tc.Err, res.Err)
			require.Equal(t, http.StatusBadRequest, rr.Code)
		})
	}
}

func TestNewHandler_Success(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		event *model.Event
	}{
		{
			event: &model.Event{
				EventType:  gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run("sucsess", func(t *testing.T) {
			t.Parallel()
			metricMocks := mocks.NewMetricService(t)
			handler := New(ctx, metricMocks)
			b, err := json.Marshal(tc.event)
			assert.NoError(t, err)
			metricMocks.On("Create",
				mock.Anything,
				mock.AnythingOfType("*model.Event")).
				Return(nil)
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/track", bytes.NewReader(b))
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			require.Equal(t, http.StatusCreated, rr.Code)
		})
	}
}

func TestNewHandler_EmptyJson(t *testing.T) {
	t.Run("empty json", func(t *testing.T) {
		ctx := context.Background()
		t.Parallel()
		metricMocks := mocks.NewMetricService(t)
		handler := New(ctx, metricMocks)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/track", bytes.NewReader([]byte("")))
		if err != nil {
			assert.Error(t, err)
		}
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		var res ResponseTrack
		_ = json.NewDecoder(rr.Body).Decode(&res)
		require.Equal(t, false, res.OK)
		require.Equal(t, "expected json object", res.Err)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestNewHandler_ErrService(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		event *model.Event
		OK    bool
		err   string
	}{
		{
			event: &model.Event{
				EventType:  gofakeit.BeerName(),
				EventTime:  gofakeit.Date(),
				ScreenName: gofakeit.BeerName(),
				Action:     gofakeit.BeerName(),
			},
			OK:  false,
			err: "example error message",
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run("errors Service", func(t *testing.T) {
			metricMocks := mocks.NewMetricService(t)
			handler := New(ctx, metricMocks)
			b, err := json.Marshal(tc.event)
			if err != nil {
				assert.NoError(t, err)
			}
			metricMocks.On("Create",
				mock.Anything,
				mock.AnythingOfType("*model.Event")).
				Return(errors.New("example error message"))

			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/track", bytes.NewReader(b))
			if err != nil {
				assert.NoError(t, err)
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusInternalServerError, rr.Code)

			var respBody ResponseTrack
			err = json.NewDecoder(rr.Body).Decode(&respBody)
			require.NoError(t, err)

			require.Equal(t, tc.err, respBody.Err)
		})
	}
}

func TestSendResponse_StatusOK(t *testing.T) {
	w := httptest.NewRecorder()
	sendResponse(w, true, nil, http.StatusOK)
	require.Equal(t, w.Code, http.StatusOK)
	var respBody ResponseTrack
	err := json.NewDecoder(w.Body).Decode(&respBody)
	if err != nil {
		assert.Error(t, err)
	}
	require.Equal(t, respBody.OK, true)
}

func TestSendResponse_BadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	sendResponse(w, false, errors.New("example"), http.StatusBadRequest)
	require.Equal(t, w.Code, http.StatusBadRequest)
	var respBody ResponseTrack
	err := json.NewDecoder(w.Body).Decode(&respBody)
	if err != nil {
		assert.Error(t, err)
	}
	require.Equal(t, respBody.OK, false)
}
