package nats

import (
	"github.com/jarri-abidi/gochat/messaging"
	"github.com/jarri-abidi/gochat/notifying"
)

// nats.NewMessagingPublisher
// messaging.Service needs a messaging.EventPublisher
func NewMessagingPublisher() messaging.EventPublisher {
	return nil
}

// nats.NewNotificationPublisher
// notifying.Service needs a notifying.EventPublisher
func NewNotifyingPublisher() notifying.EventPublisher {
	return nil
}
