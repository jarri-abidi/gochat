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

const messageEventSentKey = "MESSAGE_SENT_EVENT"

func (q *inmemQueue) ConsumeSentEvent(ctx context.Context) (*messaging.SentEvent, error) {
	v := <-q.msgs[messageEventSentKey]
	event, ok := v.(messaging.SentEvent)
	if !ok {
		return nil, errors.New("event type was not MessageSentEvent")
	}
	return &event, nil
}

func (q *inmemQueue) PublishSentEvent(ctx context.Context, event messaging.SentEvent) error {
	q.msgs[messageEventSentKey] <- event
	return nil
}
