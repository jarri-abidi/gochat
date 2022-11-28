package inmem

import (
	"context"
	"errors"

	"github.com/jarri-abidi/gochat/messaging"
)

var (
	_ messaging.EventPublisher = &inmemQueue{}
	_ messaging.EventConsumer  = &inmemQueue{}
)

func NewMessagingQueue() *inmemQueue {
	return &inmemQueue{make(map[string]chan interface{})}
}

type inmemQueue struct {
	msgs map[string]chan interface{}
}

const messageSentEventKey = "MESSAGE_SENT_EVENT"
const messageRelayEventKey = "RELAY_EVENT"

func (q *inmemQueue) ConsumeSentEvent(ctx context.Context) (*messaging.SentEvent, error) {
	v := <-q.msgs[messageSentEventKey]
	event, ok := v.(messaging.SentEvent)
	if !ok {
		return nil, errors.New("event type was not MessageSentEvent")
	}
	return &event, nil
}

func (q *inmemQueue) ConsumeRelayEvent(ctx context.Context) (*messaging.RelayEvent, error) {
	// TODO: implement
	return nil, nil
}

func (q *inmemQueue) PublishSentEvent(ctx context.Context, event messaging.SentEvent) error {
	q.msgs[messageSentEventKey] <- event
	return nil
}

func (q *inmemQueue) PublishRelayEvent(ctx context.Context, event messaging.RelayEvent) error {
	q.msgs[messageRelayEventKey] <- event
	return nil
}
