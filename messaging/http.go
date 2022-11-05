package messaging

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jarri-abidi/gochat"
	"github.com/pkg/errors"
)

type server struct {
	service Service
}

func NewServer(service Service) http.Handler {
	s := server{service: service}

	router := mux.NewRouter()
	router.Handle("/relay", s.relayHandler()).Methods("PUT")
	return router
}

// server
func (s server) relayHandler() http.HandlerFunc {
	type request struct {
		id         string    `json: "id"`
		recipient  string    `json: "recipient"`
		sender     string    `json: "sender"`
		in         []string  `json: "in"`
		content    []byte    `json: "content"`
		createdAt  time.Time `json: "createdAt"`
		sentAt     time.Time `json: "sentAt"`
		receivedAt time.Time `json: "receivedAt"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request :p", 400)
			return
		}

		resp, err := s.service.Relay(r.Context(), RelayRequest{
			ReceivedMessage: gochat.NewReceivedMessage(req.id, req.recipient, req.sender, req.in, req.content, req.createdAt, req.sentAt, req.receivedAt),
		})

		if err != nil {
			http.Error(w, "internal server error", 500)
			return
		}

		json.NewEncoder(w).Encode(resp)

	}
}

// client
func relay(ctx context.Context, req RelayRequest) (*RelayResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req.ReceivedMessage)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode request")
	}

	r, err := http.NewRequestWithContext(ctx, "PUT", req.Address, &buf)
	if err != nil {
		return nil, errors.Wrap(err, "could not create http request")
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, "could not send http request")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var rr RelayResponse
		if err := json.NewDecoder(resp.Body).Decode(&rr); err != nil {
			return nil, errors.Wrap(err, "could not decode response")
		}
		return &rr, nil
	default:
		var buf bytes.Buffer
		_, err := io.Copy(&buf, resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read response")
		}

		return nil, errors.Errorf("server returned status %d with message: %s", resp.StatusCode, buf.String())
	}
}
