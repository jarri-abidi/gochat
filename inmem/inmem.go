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

func (q *inmemQueue) ConsumeMessageSentEvent(ctx context.Context) (*messaging.MessageSentEvent, error) {
	v := <-q.msgs[messageEventSentKey]
	event, ok := v.(messaging.MessageSentEvent)
	if !ok {
		return nil, errors.New("event type was not MessageSentEvent")
	}
	return &event, nil
}

func (q *inmemQueue) PublishMessageSentEvent(ctx context.Context, event messaging.MessageSentEvent) error {
	q.msgs[messageEventSentKey] <- event
	return nil
}
