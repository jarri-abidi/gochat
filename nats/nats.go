package nats

import (
	"context"

	"github.com/jarri-abidi/gochat/messaging"
)

type Publisher struct{}

func (p *Publisher) PublishSentEvent(ctx context.Context, event messaging.SentEvent) error {
	return nil
}

type Consumer struct{}

func (c *Consumer) ConsumeSentEvent(ctx context.Context) (*messaging.SentEvent, error) {
	return nil, nil
}
